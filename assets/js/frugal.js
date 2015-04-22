$(function() {
  var engine = new Bloodhound({
    local: [{value: 'red'}, {value: 'blue'}, {value: 'green'} , {value: 'yellow'}, {value: 'violet'}, {value: 'brown'}, {value: 'purple'}, {value: 'black'}, {value: 'white'}],
    datumTokenizer: function(d) {
      return Bloodhound.tokenizers.whitespace(d.value);
    },
    queryTokenizer: Bloodhound.tokenizers.whitespace
  });

  engine.initialize();

  $('#tag').tokenfield({
    typeahead: [null, { source: engine.ttAdapter() }]
  });
  
  $('#add').on('click', function() {
    console.log($('#tag').val());
  });
  
})
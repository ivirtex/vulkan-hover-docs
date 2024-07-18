# VkSampleLocationsInfoEXT(3) Manual Page

## Name

VkSampleLocationsInfoEXT - Structure specifying a set of sample
locations



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSampleLocationsInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkSampleLocationsInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkSampleCountFlagBits         sampleLocationsPerPixel;
    VkExtent2D                    sampleLocationGridSize;
    uint32_t                      sampleLocationsCount;
    const VkSampleLocationEXT*    pSampleLocations;
} VkSampleLocationsInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sampleLocationsPerPixel` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value specifying
  the number of sample locations per pixel.

- `sampleLocationGridSize` is the size of the sample location grid to
  select custom sample locations for.

- `sampleLocationsCount` is the number of sample locations in
  `pSampleLocations`.

- `pSampleLocations` is a pointer to an array of `sampleLocationsCount`
  [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationEXT.html) structures.

## <a href="#_description" class="anchor"></a>Description

This structure **can** be used either to specify the sample locations to
be used for rendering or to specify the set of sample locations an image
subresource has been last rendered with for the purposes of layout
transitions of depth/stencil images created with
`VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT`.

The sample locations in `pSampleLocations` specify
`sampleLocationsPerPixel` number of sample locations for each pixel in
the grid of the size specified in `sampleLocationGridSize`. The sample
location for sample i at the pixel grid location (x,y) is taken from
`pSampleLocations`\[(x + y × `sampleLocationGridSize.width`) ×
`sampleLocationsPerPixel` + i\].

If the render pass has a fragment density map, the implementation will
choose the sample locations for the fragment and the contents of
`pSampleLocations` **may** be ignored.

Valid Usage

- <a href="#VUID-VkSampleLocationsInfoEXT-sampleLocationsPerPixel-01526"
  id="VUID-VkSampleLocationsInfoEXT-sampleLocationsPerPixel-01526"></a>
  VUID-VkSampleLocationsInfoEXT-sampleLocationsPerPixel-01526  
  `sampleLocationsPerPixel` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in
  [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`sampleLocationSampleCounts`

- <a href="#VUID-VkSampleLocationsInfoEXT-sampleLocationsCount-01527"
  id="VUID-VkSampleLocationsInfoEXT-sampleLocationsCount-01527"></a>
  VUID-VkSampleLocationsInfoEXT-sampleLocationsCount-01527  
  `sampleLocationsCount` **must** equal `sampleLocationsPerPixel` ×
  `sampleLocationGridSize.width` × `sampleLocationGridSize.height`

Valid Usage (Implicit)

- <a href="#VUID-VkSampleLocationsInfoEXT-sType-sType"
  id="VUID-VkSampleLocationsInfoEXT-sType-sType"></a>
  VUID-VkSampleLocationsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT`

- <a href="#VUID-VkSampleLocationsInfoEXT-pSampleLocations-parameter"
  id="VUID-VkSampleLocationsInfoEXT-pSampleLocations-parameter"></a>
  VUID-VkSampleLocationsInfoEXT-pSampleLocations-parameter  
  If `sampleLocationsCount` is not `0`, `pSampleLocations` **must** be a
  valid pointer to an array of `sampleLocationsCount`
  [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationEXT.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleLocationsEXT.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassSampleLocationsEXT.html),
[vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSampleLocationsInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

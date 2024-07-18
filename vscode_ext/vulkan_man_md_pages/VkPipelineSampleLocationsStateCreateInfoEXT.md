# VkPipelineSampleLocationsStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineSampleLocationsStateCreateInfoEXT - Structure specifying
sample locations for a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

Applications **can** also control the sample locations used for
rasterization.

If the `pNext` chain of the
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
structure specified at pipeline creation time includes a
`VkPipelineSampleLocationsStateCreateInfoEXT` structure, then that
structure controls the sample locations used when rasterizing primitives
with the pipeline.

The `VkPipelineSampleLocationsStateCreateInfoEXT` structure is defined
as:

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkPipelineSampleLocationsStateCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkBool32                    sampleLocationsEnable;
    VkSampleLocationsInfoEXT    sampleLocationsInfo;
} VkPipelineSampleLocationsStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sampleLocationsEnable` controls whether custom sample locations are
  used. If `sampleLocationsEnable` is `VK_FALSE`, the default sample
  locations are used and the values specified in `sampleLocationsInfo`
  are ignored.

- `sampleLocationsInfo` is the sample locations to use during
  rasterization if `sampleLocationsEnable` is `VK_TRUE` and the graphics
  pipeline is not created with `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sampleLocationsInfo-parameter"
  id="VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sampleLocationsInfo-parameter"></a>
  VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sampleLocationsInfo-parameter  
  `sampleLocationsInfo` **must** be a valid
  [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineSampleLocationsStateCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# VkPipelineSampleLocationsStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineSampleLocationsStateCreateInfoEXT - Structure specifying sample locations for a pipeline



## [](#_c_specification)C Specification

Applications **can** also control the sample locations used for rasterization.

If the `pNext` chain of the [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html) structure specified at pipeline creation time includes a `VkPipelineSampleLocationsStateCreateInfoEXT` structure, then that structure controls the sample locations used when rasterizing primitives with the pipeline.

The `VkPipelineSampleLocationsStateCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_sample_locations
typedef struct VkPipelineSampleLocationsStateCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkBool32                    sampleLocationsEnable;
    VkSampleLocationsInfoEXT    sampleLocationsInfo;
} VkPipelineSampleLocationsStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sampleLocationsEnable` controls whether custom sample locations are used. If `sampleLocationsEnable` is `VK_FALSE`, the default sample locations are used and the values specified in `sampleLocationsInfo` are ignored.
- `sampleLocationsInfo` is the sample locations to use during rasterization if `sampleLocationsEnable` is `VK_TRUE` and the graphics pipeline is not created with `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sType-sType)VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sampleLocationsInfo-parameter)VUID-VkPipelineSampleLocationsStateCreateInfoEXT-sampleLocationsInfo-parameter  
  `sampleLocationsInfo` **must** be a valid [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineSampleLocationsStateCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
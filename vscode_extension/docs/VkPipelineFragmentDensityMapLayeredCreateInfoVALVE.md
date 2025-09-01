# VkPipelineFragmentDensityMapLayeredCreateInfoVALVE(3) Manual Page

## Name

VkPipelineFragmentDensityMapLayeredCreateInfoVALVE - Structure specifying layered fragment density map info



## [](#_c_specification)C Specification

The `VkPipelineFragmentDensityMapLayeredCreateInfoVALVE` structure is defined as:

```c++
// Provided by VK_VALVE_fragment_density_map_layered
typedef struct VkPipelineFragmentDensityMapLayeredCreateInfoVALVE {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maxFragmentDensityMapLayers;
} VkPipelineFragmentDensityMapLayeredCreateInfoVALVE;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxFragmentDensityMapLayers` is the maximum number of layers which can be used with a fragment density map.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineFragmentDensityMapLayeredCreateInfoVALVE-maxFragmentDensityMapLayers-10825)VUID-VkPipelineFragmentDensityMapLayeredCreateInfoVALVE-maxFragmentDensityMapLayers-10825  
  `maxFragmentDensityMapLayers` **must** be less than or equal to [`maxFragmentDensityMapLayers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentDensityMapLayers)

Valid Usage (Implicit)

- [](#VUID-VkPipelineFragmentDensityMapLayeredCreateInfoVALVE-sType-sType)VUID-VkPipelineFragmentDensityMapLayeredCreateInfoVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_DENSITY_MAP_LAYERED_CREATE_INFO_VALVE`

## [](#_see_also)See Also

[VK\_VALVE\_fragment\_density\_map\_layered](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_fragment_density_map_layered.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineFragmentDensityMapLayeredCreateInfoVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
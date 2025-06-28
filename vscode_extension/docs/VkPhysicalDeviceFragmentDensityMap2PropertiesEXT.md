# VkPhysicalDeviceFragmentDensityMap2PropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMap2PropertiesEXT - Structure describing additional fragment density map properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentDensityMap2PropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map2
typedef struct VkPhysicalDeviceFragmentDensityMap2PropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           subsampledLoads;
    VkBool32           subsampledCoarseReconstructionEarlyAccess;
    uint32_t           maxSubsampledArrayLayers;
    uint32_t           maxDescriptorSetSubsampledSamplers;
} VkPhysicalDeviceFragmentDensityMap2PropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`subsampledLoads` specifies if performing image data read with load operations on subsampled attachments will be resampled to the fragment density of the render pass
- []()`subsampledCoarseReconstructionEarlyAccess` specifies if performing image data read with samplers created with `flags` containing `VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT` in fragment shader will trigger additional reads during `VK_PIPELINE_STAGE_VERTEX_SHADER_BIT`
- []()`maxSubsampledArrayLayers` is the maximum number of [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) array layers for usages supporting subsampled samplers
- []()`maxDescriptorSetSubsampledSamplers` is the maximum number of subsampled samplers that **can** be included in a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html)

## [](#_description)Description

If the `VkPhysicalDeviceFragmentDensityMap2PropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentDensityMap2PropertiesEXT-sType-sType)VUID-VkPhysicalDeviceFragmentDensityMap2PropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentDensityMap2PropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
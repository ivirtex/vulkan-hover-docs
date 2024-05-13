# VkPhysicalDeviceFragmentDensityMap2PropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMap2PropertiesEXT - Structure describing
additional fragment density map properties that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentDensityMap2PropertiesEXT` structure is
defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-subsampledLoads"></span> `subsampledLoads` specifies
  if performing image data read with load operations on subsampled
  attachments will be resampled to the fragment density of the render
  pass

- <span id="limits-subsampledCoarseReconstructionEarlyAccess"></span>
  `subsampledCoarseReconstructionEarlyAccess` specifies if performing
  image data read with samplers created with `flags` containing
  `VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT` in
  fragment shader will trigger additional reads during
  `VK_PIPELINE_STAGE_VERTEX_SHADER_BIT`

- <span id="limits-maxSubsampledArrayLayers"></span>
  `maxSubsampledArrayLayers` is the maximum number of
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) array layers for usages supporting
  subsampled samplers

- <span id="limits-maxDescriptorSetSubsampledSamplers"></span>
  `maxDescriptorSetSubsampledSamplers` is the maximum number of
  subsampled samplers that **can** be included in a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentDensityMap2PropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentDensityMap2PropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentDensityMap2PropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentDensityMap2PropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_fragment_density_map2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map2.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentDensityMap2PropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

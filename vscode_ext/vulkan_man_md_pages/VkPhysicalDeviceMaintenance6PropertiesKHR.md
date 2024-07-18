# VkPhysicalDeviceMaintenance6PropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance6PropertiesKHR - Structure describing various
implementation-defined properties introduced with VK_KHR_maintenance6



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance6PropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6
typedef struct VkPhysicalDeviceMaintenance6PropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           blockTexelViewCompatibleMultipleLayers;
    uint32_t           maxCombinedImageSamplerDescriptorCount;
    VkBool32           fragmentShadingRateClampCombinerInputs;
} VkPhysicalDeviceMaintenance6PropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `blockTexelViewCompatibleMultipleLayers` is a boolean value indicating
  that an implementation supports creating image views with
  `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` where the
  `layerCount` member of `subresourceRange` is greater than `1`.

- `maxCombinedImageSamplerDescriptorCount` is the maximum number of
  combined image sampler descriptors that the implementation uses to
  access any of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">formats that require a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a> supported by the
  implementation.

- `fragmentShadingRateClampCombinerInputs` is a boolean value indicating
  that an implementation clamps the inputs to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-combining"
  target="_blank" rel="noopener">combiner operations</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMaintenance6PropertiesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance6PropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance6PropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance6PropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance6PropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# VkPhysicalDeviceMaintenance6Properties(3) Manual Page

## Name

VkPhysicalDeviceMaintenance6Properties - Structure describing various implementation-defined properties introduced with VK\_KHR\_maintenance6



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance6Properties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceMaintenance6Properties {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           blockTexelViewCompatibleMultipleLayers;
    uint32_t           maxCombinedImageSamplerDescriptorCount;
    VkBool32           fragmentShadingRateClampCombinerInputs;
} VkPhysicalDeviceMaintenance6Properties;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance6
typedef VkPhysicalDeviceMaintenance6Properties VkPhysicalDeviceMaintenance6PropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `blockTexelViewCompatibleMultipleLayers` is a boolean value indicating that an implementation supports creating image views with `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` where the `layerCount` member of `subresourceRange` is greater than `1`.
- `maxCombinedImageSamplerDescriptorCount` is the maximum number of combined image sampler descriptors that the implementation uses to access any of the [formats that require a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion) supported by the implementation.
- `fragmentShadingRateClampCombinerInputs` is a boolean value indicating that an implementation clamps the inputs to [combiner operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-combining).

If the `VkPhysicalDeviceMaintenance6Properties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance6Properties-sType-sType)VUID-VkPhysicalDeviceMaintenance6Properties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_PROPERTIES`

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance6Properties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
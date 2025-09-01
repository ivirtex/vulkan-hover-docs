# VkPhysicalDeviceMaintenance6Features(3) Manual Page

## Name

VkPhysicalDeviceMaintenance6Features - Structure describing whether the implementation supports maintenance6 functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance6Features` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceMaintenance6Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance6;
} VkPhysicalDeviceMaintenance6Features;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance6
typedef VkPhysicalDeviceMaintenance6Features VkPhysicalDeviceMaintenance6FeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`maintenance6` indicates that the implementation supports the following:
  
  - [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **can** be used when binding an index buffer
  - [VkBindMemoryStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatus.html) **can** be included in the `pNext` chain of the [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html) and [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html) structures, enabling applications to retrieve [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) values for individual memory binding operations.
  - [VkPhysicalDeviceMaintenance6Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6Properties.html)::`blockTexelViewCompatibleMultipleLayers` property to indicate that the implementation supports creating image views with `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` where the `layerCount` member of `subresourceRange` is greater than `1`.
  - [VkPhysicalDeviceMaintenance6Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6Properties.html)::`maxCombinedImageSamplerDescriptorCount` property which indicates the maximum descriptor size required for any [format that requires a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion) supported by the implementation.
  - A [VkPhysicalDeviceMaintenance6Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6Properties.html)::`fragmentShadingRateClampCombinerInputs` property which indicates whether the implementation clamps the inputs to fragment shading rate combiner operations.

If the `VkPhysicalDeviceMaintenance6Features` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMaintenance6Features`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance6Features-sType-sType)VUID-VkPhysicalDeviceMaintenance6Features-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_FEATURES`

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance6Features).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
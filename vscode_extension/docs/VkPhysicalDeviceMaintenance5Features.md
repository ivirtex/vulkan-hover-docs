# VkPhysicalDeviceMaintenance5Features(3) Manual Page

## Name

VkPhysicalDeviceMaintenance5Features - Structure describing whether the implementation supports maintenance5 functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance5Features` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceMaintenance5Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance5;
} VkPhysicalDeviceMaintenance5Features;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkPhysicalDeviceMaintenance5Features VkPhysicalDeviceMaintenance5FeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`maintenance5` indicates that the implementation supports the following:
  
  - The ability to expose support for the optional format `VK_FORMAT_A1B5G5R5_UNORM_PACK16`.
  - The ability to expose support for the optional format `VK_FORMAT_A8_UNORM`.
  - A property to indicate that multisample coverage operations are performed after sample counting in EarlyFragmentTests mode.
  - Creating a `VkBufferView` with a subset of the associated `VkBuffer` usage using [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html).
  - A new function [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html), allowing a range of memory to be bound as an index buffer.
  - [vkGetDeviceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceProcAddr.html) will return `NULL` for function pointers of core functions for versions higher than the version requested by the application.
  - [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers2.html) supports using `VK_WHOLE_SIZE` in the `pSizes` parameter.
  - If `PointSize` is not written, a default value of `1.0` is used for the size of points.
  - [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) **can** be added as a chained structure to pipeline creation via [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), rather than having to create a shader module.
  - A function [vkGetRenderingAreaGranularity](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRenderingAreaGranularity.html) to query the optimal render area for a dynamic rendering instance.
  - A property to indicate that depth/stencil texturing operations with `VK_COMPONENT_SWIZZLE_ONE` have defined behavior.
  - [vkGetDeviceImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayout.html) allows an application to perform a [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout.html) query without having to create an image.
  - `VK_REMAINING_ARRAY_LAYERS` as the `layerCount` member of [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html).
  - A property to indicate whether `PointSize` controls the final rasterization of polygons if [polygon mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode) is `VK_POLYGON_MODE_POINT`.
  - Two properties to indicate the non-strict line rasterization algorithm used.
  - Two new flags words [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html) and [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html).
  - Physical-device-level functions **can** now be called with any value in the valid range for a type beyond the defined enumerants, such that applications can avoid checking individual features, extensions, or versions before querying supported properties of a particular enumerant.
  - Copies between images of any type are allowed, with 1D images treated as 2D images with a height of `1`.

If the `VkPhysicalDeviceMaintenance5Features` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMaintenance5Features`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance5Features-sType-sType)VUID-VkPhysicalDeviceMaintenance5Features-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_FEATURES`

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance5Features)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
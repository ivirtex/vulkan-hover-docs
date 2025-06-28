# VkPhysicalDeviceDynamicRenderingLocalReadFeatures(3) Manual Page

## Name

VkPhysicalDeviceDynamicRenderingLocalReadFeatures - Structure indicating support for local reads in dynamic render pass instances



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDynamicRenderingLocalReadFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceDynamicRenderingLocalReadFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           dynamicRenderingLocalRead;
} VkPhysicalDeviceDynamicRenderingLocalReadFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_dynamic_rendering_local_read
typedef VkPhysicalDeviceDynamicRenderingLocalReadFeatures VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`dynamicRenderingLocalRead` specifies that the implementation supports local reads inside dynamic render pass instances using the [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) command.

If the `VkPhysicalDeviceDynamicRenderingLocalReadFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDynamicRenderingLocalReadFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDynamicRenderingLocalReadFeatures-sType-sType)VUID-VkPhysicalDeviceDynamicRenderingLocalReadFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_LOCAL_READ_FEATURES`

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDynamicRenderingLocalReadFeatures)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
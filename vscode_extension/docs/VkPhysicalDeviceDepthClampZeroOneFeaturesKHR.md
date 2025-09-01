# VkPhysicalDeviceDepthClampZeroOneFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceDepthClampZeroOneFeaturesKHR - Structure describing feature to control zero to one depth clamping



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDepthClampZeroOneFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_depth_clamp_zero_one
typedef struct VkPhysicalDeviceDepthClampZeroOneFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           depthClampZeroOne;
} VkPhysicalDeviceDepthClampZeroOneFeaturesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_depth_clamp_zero_one
typedef VkPhysicalDeviceDepthClampZeroOneFeaturesKHR VkPhysicalDeviceDepthClampZeroOneFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`depthClampZeroOne` indicates that the implementation supports clamping the depth to a range of `0` to `1`.

If the `VkPhysicalDeviceDepthClampZeroOneFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDepthClampZeroOneFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDepthClampZeroOneFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceDepthClampZeroOneFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLAMP_ZERO_ONE_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_EXT\_depth\_clamp\_zero\_one](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_zero_one.html), [VK\_KHR\_depth\_clamp\_zero\_one](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_clamp_zero_one.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDepthClampZeroOneFeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
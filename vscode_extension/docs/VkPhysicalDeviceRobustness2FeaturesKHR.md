# VkPhysicalDeviceRobustness2FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceRobustness2FeaturesKHR - Structure describing the out-of-bounds behavior for an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRobustness2FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_robustness2
typedef struct VkPhysicalDeviceRobustness2FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           robustBufferAccess2;
    VkBool32           robustImageAccess2;
    VkBool32           nullDescriptor;
} VkPhysicalDeviceRobustness2FeaturesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_robustness2
typedef VkPhysicalDeviceRobustness2FeaturesKHR VkPhysicalDeviceRobustness2FeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`robustBufferAccess2` enables [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-buffer-access2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-buffer-access2) guarantees for shader buffer accesses.
- []()`robustImageAccess2` enables [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-image-access2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-image-access2) guarantees for shader image accesses.
- []()`nullDescriptor` indicates whether descriptors **can** be written with a [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) resource or view, which are considered valid to access and act as if the descriptor were bound to nothing.

## [](#_description)Description

If the `VkPhysicalDeviceRobustness2FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRobustness2FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage

- [](#VUID-VkPhysicalDeviceRobustness2FeaturesKHR-robustBufferAccess2-04000)VUID-VkPhysicalDeviceRobustness2FeaturesKHR-robustBufferAccess2-04000  
  If `robustBufferAccess2` is enabled then [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) **must** also be enabled

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRobustness2FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceRobustness2FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_robustness2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_robustness2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRobustness2FeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
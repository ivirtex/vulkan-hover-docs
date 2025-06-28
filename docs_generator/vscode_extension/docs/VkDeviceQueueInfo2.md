# VkDeviceQueueInfo2(3) Manual Page

## Name

VkDeviceQueueInfo2 - Structure specifying the parameters used for device queue creation



## [](#_c_specification)C Specification

The `VkDeviceQueueInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDeviceQueueInfo2 {
    VkStructureType             sType;
    const void*                 pNext;
    VkDeviceQueueCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
    uint32_t                    queueIndex;
} VkDeviceQueueInfo2;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure. The `pNext` chain of `VkDeviceQueueInfo2` **can** be used to provide additional device queue parameters to `vkGetDeviceQueue2`.
- `flags` is a [VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateFlags.html) value indicating the flags used to create the device queue.
- `queueFamilyIndex` is the index of the queue family to which the queue belongs.
- `queueIndex` is the index of the queue to retrieve from within the set of queues that share both the queue family and flags specified.

## [](#_description)Description

The queue returned by `vkGetDeviceQueue2` **must** have the same `flags` value from this structure as that used at device creation time in a [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structure.

Note

Normally, if you create both protected-capable and non-protected-capable queues with the same family, they are treated as separate lists of queues and `queueIndex` is relative to the start of the list of queues specified by both `queueFamilyIndex` and `flags`. However, for historical reasons, some implementations may exhibit different behavior. These divergent implementations instead concatenate the lists of queues and treat `queueIndex` as relative to the start of the first list of queues with the given `queueFamilyIndex`. This only matters in cases where an application has created both protected-capable and non-protected-capable queues from the same queue family.

For such divergent implementations, the maximum value of `queueIndex` is equal to the sum of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html)::`queueCount` minus one, for all [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structures that share a common `queueFamilyIndex`.

Such implementations will return `NULL` for either the protected or unprotected queues when calling `vkGetDeviceQueue2` with `queueIndex` in the range zero to [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html)::`queueCount` minus one. In cases where these implementations returned `NULL`, the corresponding queues are instead located in the extended range described in the preceding two paragraphs.

This behavior will not be observed on any driver that has passed Vulkan conformance test suite version 1.3.3.0, or any subsequent version. This information can be found by querying `VkPhysicalDeviceDriverProperties`::`conformanceVersion`.

Valid Usage

- [](#VUID-VkDeviceQueueInfo2-queueFamilyIndex-01842)VUID-VkDeviceQueueInfo2-queueFamilyIndex-01842  
  `queueFamilyIndex` **must** be one of the queue family indices specified when `device` was created, via the [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structure
- [](#VUID-VkDeviceQueueInfo2-flags-06225)VUID-VkDeviceQueueInfo2-flags-06225  
  `flags` **must** be equal to [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html)::`flags` for a [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structure for the queue family indicated by `queueFamilyIndex` when `device` was created
- [](#VUID-VkDeviceQueueInfo2-queueIndex-01843)VUID-VkDeviceQueueInfo2-queueIndex-01843  
  `queueIndex` **must** be less than [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html)::`queueCount` for the corresponding queue family and flags indicated by `queueFamilyIndex` and `flags` when `device` was created

Valid Usage (Implicit)

- [](#VUID-VkDeviceQueueInfo2-sType-sType)VUID-VkDeviceQueueInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2`
- [](#VUID-VkDeviceQueueInfo2-pNext-pNext)VUID-VkDeviceQueueInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDeviceQueueInfo2-flags-parameter)VUID-VkDeviceQueueInfo2-flags-parameter  
  `flags` **must** be a valid combination of [VkDeviceQueueCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceQueue2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceQueue2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceQueueInfo2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
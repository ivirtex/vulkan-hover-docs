# VkPhysicalDeviceExternalBufferInfo(3) Manual Page

## Name

VkPhysicalDeviceExternalBufferInfo - Structure specifying buffer creation parameters



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExternalBufferInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceExternalBufferInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkBufferCreateFlags                   flags;
    VkBufferUsageFlags                    usage;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkPhysicalDeviceExternalBufferInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory_capabilities
typedef VkPhysicalDeviceExternalBufferInfo VkPhysicalDeviceExternalBufferInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html) describing additional parameters of the buffer, corresponding to [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`flags`.
- `usage` is a bitmask of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) describing the intended usage of the buffer, corresponding to [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`usage`.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the memory handle type that will be used with the memory associated with the buffer.

## [](#_description)Description

Only usage flags representable in [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) are returned in this structure’s `usage`. If the `pNext` chain includes a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, all usage flags of the buffer are returned in [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html)::`usage`.

Valid Usage

- [](#VUID-VkPhysicalDeviceExternalBufferInfo-None-09499)VUID-VkPhysicalDeviceExternalBufferInfo-None-09499  
  If the `pNext` chain does not include a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, `usage` **must** be a valid combination of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) values
- [](#VUID-VkPhysicalDeviceExternalBufferInfo-None-09500)VUID-VkPhysicalDeviceExternalBufferInfo-None-09500  
  If the `pNext` chain does not include a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, `usage` **must** not be 0

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalBufferInfo-sType-sType)VUID-VkPhysicalDeviceExternalBufferInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO`
- [](#VUID-VkPhysicalDeviceExternalBufferInfo-pNext-pNext)VUID-VkPhysicalDeviceExternalBufferInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html)
- [](#VUID-VkPhysicalDeviceExternalBufferInfo-sType-unique)VUID-VkPhysicalDeviceExternalBufferInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPhysicalDeviceExternalBufferInfo-flags-parameter)VUID-VkPhysicalDeviceExternalBufferInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html) values
- [](#VUID-VkPhysicalDeviceExternalBufferInfo-handleType-parameter)VUID-VkPhysicalDeviceExternalBufferInfo-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBufferCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlags.html), [VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalBufferProperties.html), [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalBufferProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalBufferInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
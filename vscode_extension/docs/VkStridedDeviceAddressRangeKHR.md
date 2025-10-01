# VkStridedDeviceAddressRangeKHR(3) Manual Page

## Name

VkStridedDeviceAddressRangeKHR - Structure specifying a device address range with a stride



## [](#_c_specification)C Specification

A strided device address range is defined by the structure:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkStridedDeviceAddressRangeKHR {
    VkDeviceAddress    address;
    VkDeviceSize       size;
    VkDeviceSize       stride;
} VkStridedDeviceAddressRangeKHR;
```

## [](#_members)Members

- `address` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) specifying the start of the range.
- `size` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) specifying the size of the range.
- `stride` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) specifying the stride of elements over the range.

## [](#_description)Description

Valid Usage

- [](#VUID-VkStridedDeviceAddressRangeKHR-stride-10957)VUID-VkStridedDeviceAddressRangeKHR-stride-10957  
  `stride` **must** be less than or equal to `size`
- [](#VUID-VkStridedDeviceAddressRangeKHR-address-11365)VUID-VkStridedDeviceAddressRangeKHR-address-11365  
  The sum of `address` and `size` **must** be less than or equal to the sum of an address retrieved from a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) and the value of [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`size` used to create that [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html)
- [](#VUID-VkStridedDeviceAddressRangeKHR-size-11366)VUID-VkStridedDeviceAddressRangeKHR-size-11366  
  If `size` is not 0 and the buffer from which `address` was queried is non-sparse, then the buffer **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object

Valid Usage (Implicit)

- [](#VUID-VkStridedDeviceAddressRangeKHR-address-parameter)VUID-VkStridedDeviceAddressRangeKHR-address-parameter  
  `address` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkCopyMemoryIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectInfoKHR.html), [VkCopyMemoryToImageIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectInfoKHR.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkStridedDeviceAddressRangeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkCopyMemoryIndirectCommandKHR(3) Manual Page

## Name

VkCopyMemoryIndirectCommandKHR - Structure specifying indirect memory region copy operation



## [](#_c_specification)C Specification

The structure describing source and destination memory regions, `VkCopyMemoryIndirectCommandKHR` is defined as:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkCopyMemoryIndirectCommandKHR {
    VkDeviceAddress    srcAddress;
    VkDeviceAddress    dstAddress;
    VkDeviceSize       size;
} VkCopyMemoryIndirectCommandKHR;
```

or the equivalent

```c++
// Provided by VK_NV_copy_memory_indirect
typedef VkCopyMemoryIndirectCommandKHR VkCopyMemoryIndirectCommandNV;
```

## [](#_members)Members

- `srcAddress` is the starting address of the source device memory to copy from.
- `dstAddress` is the starting address of the destination device memory to copy to.
- `size` is the size of the copy in bytes.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryIndirectCommandKHR-srcAddress-10958)VUID-VkCopyMemoryIndirectCommandKHR-srcAddress-10958  
  The `srcAddress` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryIndirectCommandKHR-dstAddress-10959)VUID-VkCopyMemoryIndirectCommandKHR-dstAddress-10959  
  The `dstAddress` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryIndirectCommandKHR-size-10960)VUID-VkCopyMemoryIndirectCommandKHR-size-10960  
  The `size` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryIndirectCommandKHR-srcAddress-10961)VUID-VkCopyMemoryIndirectCommandKHR-srcAddress-10961  
  The memory in range [`srcAddress`, `srcAddress`  
  `size` - 1] **must** be within the bounds of the memory allocation backing `srcAddress`
- [](#VUID-VkCopyMemoryIndirectCommandKHR-dstAddress-10962)VUID-VkCopyMemoryIndirectCommandKHR-dstAddress-10962  
  The memory in range [`dstAddress`, `dstAddress`  
  `size` - 1] **must** be within the bounds of the memory allocation backing `dstAddress`

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryIndirectCommandKHR-srcAddress-parameter)VUID-VkCopyMemoryIndirectCommandKHR-srcAddress-parameter  
  `srcAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkCopyMemoryIndirectCommandKHR-dstAddress-parameter)VUID-VkCopyMemoryIndirectCommandKHR-dstAddress-parameter  
  `dstAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VK\_NV\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_copy_memory_indirect.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryIndirectCommandKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
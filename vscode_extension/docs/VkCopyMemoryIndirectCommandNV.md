# VkCopyMemoryIndirectCommandNV(3) Manual Page

## Name

VkCopyMemoryIndirectCommandNV - Structure specifying indirect memory region copy operation



## [](#_c_specification)C Specification

The structure describing source and destination memory regions, `VkCopyMemoryIndirectCommandNV` is defined as:

```c++
// Provided by VK_NV_copy_memory_indirect
typedef struct VkCopyMemoryIndirectCommandNV {
    VkDeviceAddress    srcAddress;
    VkDeviceAddress    dstAddress;
    VkDeviceSize       size;
} VkCopyMemoryIndirectCommandNV;
```

## [](#_members)Members

- `srcAddress` is the starting address of the source device memory to copy from.
- `dstAddress` is the starting address of the destination device memory to copy to.
- `size` is the size of the copy in bytes.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryIndirectCommandNV-srcAddress-07657)VUID-VkCopyMemoryIndirectCommandNV-srcAddress-07657  
  The `srcAddress` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryIndirectCommandNV-dstAddress-07658)VUID-VkCopyMemoryIndirectCommandNV-dstAddress-07658  
  The `dstAddress` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryIndirectCommandNV-size-07659)VUID-VkCopyMemoryIndirectCommandNV-size-07659  
  The `size` **must** be 4 byte aligned

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryIndirectCommandNV-srcAddress-parameter)VUID-VkCopyMemoryIndirectCommandNV-srcAddress-parameter  
  `srcAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkCopyMemoryIndirectCommandNV-dstAddress-parameter)VUID-VkCopyMemoryIndirectCommandNV-dstAddress-parameter  
  `dstAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_NV\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_copy_memory_indirect.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
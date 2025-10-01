# VkCopyMemoryIndirectInfoKHR(3) Manual Page

## Name

VkCopyMemoryIndirectInfoKHR - Parameters describing indirect copy parameters



## [](#_c_specification)C Specification

The [VkCopyMemoryIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkCopyMemoryIndirectInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkAddressCopyFlagsKHR             srcCopyFlags;
    VkAddressCopyFlagsKHR             dstCopyFlags;
    uint32_t                          copyCount;
    VkStridedDeviceAddressRangeKHR    copyAddressRange;
} VkCopyMemoryIndirectInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcCopyFlags` is a [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html) value defining the copy flags for the source address range.
- `dstCopyFlags` is a [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html) value defining the copy flags for the destination address range.
- `copyCount` is the number of copies to execute, and **can** be zero.
- `copyAddressRange` is a memory region specifying the copy parameters. It is laid out as an array of [VkCopyMemoryIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandKHR.html) structures.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-10938)VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-10938  
  If `srcCopyFlags` contains `VK_ADDRESS_COPY_SPARSE_BIT_KHR`, the source memory regions accessed **must** be fully-resident
- [](#VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-10939)VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-10939  
  If `dstCopyFlags` contains `VK_ADDRESS_COPY_SPARSE_BIT_KHR`, the destination memory regions accessed **must** be fully-resident
- [](#VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-10940)VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-10940  
  `srcCopyFlags` **must** not contain `VK_ADDRESS_COPY_PROTECTED_BIT_KHR`
- [](#VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-10941)VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-10941  
  `dstCopyFlags` **must** not contain `VK_ADDRESS_COPY_PROTECTED_BIT_KHR`
- [](#VUID-VkCopyMemoryIndirectInfoKHR-copyAddressRange-10942)VUID-VkCopyMemoryIndirectInfoKHR-copyAddressRange-10942  
  `copyAddressRange.pname`:address **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryIndirectInfoKHR-copyAddressRange-10943)VUID-VkCopyMemoryIndirectInfoKHR-copyAddressRange-10943  
  `copyAddressRange.pname`:stride **must** be a multiple of `4` and **must** be greater than or equal to sizeof([VkCopyMemoryIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandKHR.html))
- [](#VUID-VkCopyMemoryIndirectInfoKHR-copyCount-10944)VUID-VkCopyMemoryIndirectInfoKHR-copyCount-10944  
  `copyCount` **must** be less than or equal to `copyAddressRange.pname`:size / `copyAddressRange.pname`:stride
- [](#VUID-VkCopyMemoryIndirectInfoKHR-copyAddressRange-10945)VUID-VkCopyMemoryIndirectInfoKHR-copyAddressRange-10945  
  Any of the source or destination memory regions specified in `copyAddressRange` **must** not overlap with any of the specified destination memory regions

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryIndirectInfoKHR-sType-sType)VUID-VkCopyMemoryIndirectInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MEMORY_INDIRECT_INFO_KHR`
- [](#VUID-VkCopyMemoryIndirectInfoKHR-pNext-pNext)VUID-VkCopyMemoryIndirectInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-parameter)VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-parameter  
  `srcCopyFlags` **must** be a valid combination of [VkAddressCopyFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagBitsKHR.html) values
- [](#VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-requiredbitmask)VUID-VkCopyMemoryIndirectInfoKHR-srcCopyFlags-requiredbitmask  
  `srcCopyFlags` **must** not be `0`
- [](#VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-parameter)VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-parameter  
  `dstCopyFlags` **must** be a valid combination of [VkAddressCopyFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagBitsKHR.html) values
- [](#VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-requiredbitmask)VUID-VkCopyMemoryIndirectInfoKHR-dstCopyFlags-requiredbitmask  
  `dstCopyFlags` **must** not be `0`

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html), [VkStridedDeviceAddressRangeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRangeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyMemoryIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryIndirectKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryIndirectInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
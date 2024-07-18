# VkCopyMemoryIndirectCommandNV(3) Manual Page

## Name

VkCopyMemoryIndirectCommandNV - Structure specifying indirect memory
region copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The structure describing source and destination memory regions,
`VkCopyMemoryIndirectCommandNV` is defined as:

``` c
// Provided by VK_NV_copy_memory_indirect
typedef struct VkCopyMemoryIndirectCommandNV {
    VkDeviceAddress    srcAddress;
    VkDeviceAddress    dstAddress;
    VkDeviceSize       size;
} VkCopyMemoryIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `srcAddress` is the starting address of the source device memory to
  copy from.

- `dstAddress` is the starting address of the destination device memory
  to copy to.

- `size` is the size of the copy in bytes.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyMemoryIndirectCommandNV-srcAddress-07657"
  id="VUID-VkCopyMemoryIndirectCommandNV-srcAddress-07657"></a>
  VUID-VkCopyMemoryIndirectCommandNV-srcAddress-07657  
  The `srcAddress` **must** be 4 byte aligned

- <a href="#VUID-VkCopyMemoryIndirectCommandNV-dstAddress-07658"
  id="VUID-VkCopyMemoryIndirectCommandNV-dstAddress-07658"></a>
  VUID-VkCopyMemoryIndirectCommandNV-dstAddress-07658  
  The `dstAddress` **must** be 4 byte aligned

- <a href="#VUID-VkCopyMemoryIndirectCommandNV-size-07659"
  id="VUID-VkCopyMemoryIndirectCommandNV-size-07659"></a>
  VUID-VkCopyMemoryIndirectCommandNV-size-07659  
  The `size` **must** be 4 byte aligned

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_copy_memory_indirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_copy_memory_indirect.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMemoryIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

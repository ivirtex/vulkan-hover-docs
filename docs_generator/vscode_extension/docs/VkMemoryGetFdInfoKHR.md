# VkMemoryGetFdInfoKHR(3) Manual Page

## Name

VkMemoryGetFdInfoKHR - Structure describing a POSIX FD memory export operation



## [](#_c_specification)C Specification

The `VkMemoryGetFdInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_memory_fd
typedef struct VkMemoryGetFdInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetFdInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is the memory object from which the handle will be exported.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the file descriptor exported depend on the value of `handleType`. See [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) for a description of the properties of the defined external memory handle types.

Note

The size of the exported file **may** be larger than the size requested by [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html)::`allocationSize`. If `handleType` is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`, then the application **can** query the file’s actual size with [`lseek`](https://man7.org/linux/man-pages/man2/lseek.2.html).

Valid Usage

- [](#VUID-VkMemoryGetFdInfoKHR-handleType-00671)VUID-VkMemoryGetFdInfoKHR-handleType-00671  
  `handleType` **must** have been included in [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` when `memory` was created
- [](#VUID-VkMemoryGetFdInfoKHR-handleType-00672)VUID-VkMemoryGetFdInfoKHR-handleType-00672  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkMemoryGetFdInfoKHR-sType-sType)VUID-VkMemoryGetFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR`
- [](#VUID-VkMemoryGetFdInfoKHR-pNext-pNext)VUID-VkMemoryGetFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryGetFdInfoKHR-memory-parameter)VUID-VkMemoryGetFdInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkMemoryGetFdInfoKHR-handleType-parameter)VUID-VkMemoryGetFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_fd.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryFdKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryGetFdInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
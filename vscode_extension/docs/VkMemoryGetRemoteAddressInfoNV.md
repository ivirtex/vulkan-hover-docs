# VkMemoryGetRemoteAddressInfoNV(3) Manual Page

## Name

VkMemoryGetRemoteAddressInfoNV - Structure describing a remote accessible address export operation



## [](#_c_specification)C Specification

The `VkMemoryGetRemoteAddressInfoNV` structure is defined as:

```c++
// Provided by VK_NV_external_memory_rdma
typedef struct VkMemoryGetRemoteAddressInfoNV {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetRemoteAddressInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is the memory object from which the remote accessible address will be exported.
- `handleType` is the type of handle requested.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryGetRemoteAddressInfoNV-handleType-04966)VUID-VkMemoryGetRemoteAddressInfoNV-handleType-04966  
  `handleType` **must** have been included in [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` when `memory` was created

Valid Usage (Implicit)

- [](#VUID-VkMemoryGetRemoteAddressInfoNV-sType-sType)VUID-VkMemoryGetRemoteAddressInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_REMOTE_ADDRESS_INFO_NV`
- [](#VUID-VkMemoryGetRemoteAddressInfoNV-pNext-pNext)VUID-VkMemoryGetRemoteAddressInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryGetRemoteAddressInfoNV-memory-parameter)VUID-VkMemoryGetRemoteAddressInfoNV-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkMemoryGetRemoteAddressInfoNV-handleType-parameter)VUID-VkMemoryGetRemoteAddressInfoNV-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_rdma](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_rdma.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryRemoteAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryRemoteAddressNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryGetRemoteAddressInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
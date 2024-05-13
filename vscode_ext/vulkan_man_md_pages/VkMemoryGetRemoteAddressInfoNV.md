# VkMemoryGetRemoteAddressInfoNV(3) Manual Page

## Name

VkMemoryGetRemoteAddressInfoNV - Structure describing a remote
accessible address export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryGetRemoteAddressInfoNV` structure is defined as:

``` c
// Provided by VK_NV_external_memory_rdma
typedef struct VkMemoryGetRemoteAddressInfoNV {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetRemoteAddressInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` is the memory object from which the remote accessible address
  will be exported.

- `handleType` is the type of handle requested.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMemoryGetRemoteAddressInfoNV-handleType-04966"
  id="VUID-VkMemoryGetRemoteAddressInfoNV-handleType-04966"></a>
  VUID-VkMemoryGetRemoteAddressInfoNV-handleType-04966  
  `handleType` **must** have been included in
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  when `memory` was created

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryGetRemoteAddressInfoNV-sType-sType"
  id="VUID-VkMemoryGetRemoteAddressInfoNV-sType-sType"></a>
  VUID-VkMemoryGetRemoteAddressInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_GET_REMOTE_ADDRESS_INFO_NV`

- <a href="#VUID-VkMemoryGetRemoteAddressInfoNV-pNext-pNext"
  id="VUID-VkMemoryGetRemoteAddressInfoNV-pNext-pNext"></a>
  VUID-VkMemoryGetRemoteAddressInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryGetRemoteAddressInfoNV-memory-parameter"
  id="VUID-VkMemoryGetRemoteAddressInfoNV-memory-parameter"></a>
  VUID-VkMemoryGetRemoteAddressInfoNV-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-VkMemoryGetRemoteAddressInfoNV-handleType-parameter"
  id="VUID-VkMemoryGetRemoteAddressInfoNV-handleType-parameter"></a>
  VUID-VkMemoryGetRemoteAddressInfoNV-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_rdma](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_rdma.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryRemoteAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryRemoteAddressNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryGetRemoteAddressInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

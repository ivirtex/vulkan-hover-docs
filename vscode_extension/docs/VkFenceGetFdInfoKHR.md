# VkFenceGetFdInfoKHR(3) Manual Page

## Name

VkFenceGetFdInfoKHR - Structure describing a POSIX FD fence export
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFenceGetFdInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_fence_fd
typedef struct VkFenceGetFdInfoKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    VkFence                              fence;
    VkExternalFenceHandleTypeFlagBits    handleType;
} VkFenceGetFdInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fence` is the fence from which state will be exported.

- `handleType` is a
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value specifying the type of handle requested.

## <a href="#_description" class="anchor"></a>Description

The properties of the file descriptor returned depend on the value of
`handleType`. See
[VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
for a description of the properties of the defined external fence handle
types.

Valid Usage

- <a href="#VUID-VkFenceGetFdInfoKHR-handleType-01453"
  id="VUID-VkFenceGetFdInfoKHR-handleType-01453"></a>
  VUID-VkFenceGetFdInfoKHR-handleType-01453  
  `handleType` **must** have been included in
  [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html)::`handleTypes`
  when `fence`’s current payload was created

- <a href="#VUID-VkFenceGetFdInfoKHR-handleType-01454"
  id="VUID-VkFenceGetFdInfoKHR-handleType-01454"></a>
  VUID-VkFenceGetFdInfoKHR-handleType-01454  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `fence` **must** be signaled, or have an associated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-signaling"
  target="_blank" rel="noopener">fence signal operation</a> pending
  execution

- <a href="#VUID-VkFenceGetFdInfoKHR-fence-01455"
  id="VUID-VkFenceGetFdInfoKHR-fence-01455"></a>
  VUID-VkFenceGetFdInfoKHR-fence-01455  
  `fence` **must** not currently have its payload replaced by an
  imported payload as described below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-importing"
  target="_blank" rel="noopener">Importing Fence Payloads</a> unless
  that imported payload’s handle type was included in
  [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)::`exportFromImportedHandleTypes`
  for `handleType`

- <a href="#VUID-VkFenceGetFdInfoKHR-handleType-01456"
  id="VUID-VkFenceGetFdInfoKHR-handleType-01456"></a>
  VUID-VkFenceGetFdInfoKHR-handleType-01456  
  `handleType` **must** be defined as a POSIX file descriptor handle

Valid Usage (Implicit)

- <a href="#VUID-VkFenceGetFdInfoKHR-sType-sType"
  id="VUID-VkFenceGetFdInfoKHR-sType-sType"></a>
  VUID-VkFenceGetFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR`

- <a href="#VUID-VkFenceGetFdInfoKHR-pNext-pNext"
  id="VUID-VkFenceGetFdInfoKHR-pNext-pNext"></a>
  VUID-VkFenceGetFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkFenceGetFdInfoKHR-fence-parameter"
  id="VUID-VkFenceGetFdInfoKHR-fence-parameter"></a>
  VUID-VkFenceGetFdInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-VkFenceGetFdInfoKHR-handleType-parameter"
  id="VUID-VkFenceGetFdInfoKHR-handleType-parameter"></a>
  VUID-VkFenceGetFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_fd.html),
[VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetFenceFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceFdKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFenceGetFdInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# VkSemaphoreGetFdInfoKHR(3) Manual Page

## Name

VkSemaphoreGetFdInfoKHR - Structure describing a POSIX FD semaphore
export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreGetFdInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_semaphore_fd
typedef struct VkSemaphoreGetFdInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkSemaphoreGetFdInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `semaphore` is the semaphore from which state will be exported.

- `handleType` is a
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value specifying the type of handle requested.

## <a href="#_description" class="anchor"></a>Description

The properties of the file descriptor returned depend on the value of
`handleType`. See
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
for a description of the properties of the defined external semaphore
handle types.

Valid Usage

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-01132"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-01132"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-01132  
  `handleType` **must** have been included in
  [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html)::`handleTypes`
  when `semaphore`’s current payload was created

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-semaphore-01133"
  id="VUID-VkSemaphoreGetFdInfoKHR-semaphore-01133"></a>
  VUID-VkSemaphoreGetFdInfoKHR-semaphore-01133  
  `semaphore` **must** not currently have its payload replaced by an
  imported payload as described below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a> unless
  that imported payload’s handle type was included in
  [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)::`exportFromImportedHandleTypes`
  for `handleType`

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-01134"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-01134"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-01134  
  If `handleType` refers to a handle type with copy payload transference
  semantics, as defined below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a>, there
  **must** be no queue waiting on `semaphore`

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-01135"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-01135"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-01135  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `semaphore` **must** be signaled, or have an associated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> pending
  execution

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-01136"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-01136"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-01136  
  `handleType` **must** be defined as a POSIX file descriptor handle

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-03253"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-03253"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-03253  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `semaphore` **must** have been created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-03254"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-03254"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-03254  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `semaphore` **must** have an associated semaphore signal
  operation that has been submitted for execution and any semaphore
  signal operations on which it depends **must** have also been
  submitted for execution

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-sType-sType"
  id="VUID-VkSemaphoreGetFdInfoKHR-sType-sType"></a>
  VUID-VkSemaphoreGetFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR`

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-pNext-pNext"
  id="VUID-VkSemaphoreGetFdInfoKHR-pNext-pNext"></a>
  VUID-VkSemaphoreGetFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-semaphore-parameter"
  id="VUID-VkSemaphoreGetFdInfoKHR-semaphore-parameter"></a>
  VUID-VkSemaphoreGetFdInfoKHR-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkSemaphoreGetFdInfoKHR-handleType-parameter"
  id="VUID-VkSemaphoreGetFdInfoKHR-handleType-parameter"></a>
  VUID-VkSemaphoreGetFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_fd.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreFdKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreGetFdInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

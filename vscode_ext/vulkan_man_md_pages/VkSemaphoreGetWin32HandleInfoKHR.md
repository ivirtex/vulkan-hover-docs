# VkSemaphoreGetWin32HandleInfoKHR(3) Manual Page

## Name

VkSemaphoreGetWin32HandleInfoKHR - Structure describing a Win32 handle
semaphore export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreGetWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_semaphore_win32
typedef struct VkSemaphoreGetWin32HandleInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkSemaphoreGetWin32HandleInfoKHR;
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

The properties of the handle returned depend on the value of
`handleType`. See
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
for a description of the properties of the defined external semaphore
handle types.

Valid Usage

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01126"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01126"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01126  
  `handleType` **must** have been included in
  [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html)::`handleTypes`
  when the `semaphore`’s current payload was created

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01127"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01127"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01127  
  If `handleType` is defined as an NT handle,
  [vkGetSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreWin32HandleKHR.html)
  **must** be called no more than once for each valid unique combination
  of `semaphore` and `handleType`

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-semaphore-01128"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-semaphore-01128"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-semaphore-01128  
  `semaphore` **must** not currently have its payload replaced by an
  imported payload as described below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a> unless
  that imported payload’s handle type was included in
  [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)::`exportFromImportedHandleTypes`
  for `handleType`

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01129"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01129"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01129  
  If `handleType` refers to a handle type with copy payload transference
  semantics, as defined below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a>, there
  **must** be no queue waiting on `semaphore`

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01130"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01130"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01130  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `semaphore` **must** be signaled, or have an associated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> pending
  execution

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01131"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01131"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-01131  
  `handleType` **must** be defined as an NT handle or a global share
  handle

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-sType-sType"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-pNext-pNext"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-pNext-pNext"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-semaphore-parameter"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-semaphore-parameter"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-parameter"
  id="VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-parameter"></a>
  VUID-VkSemaphoreGetWin32HandleInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_win32.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreWin32HandleKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreGetWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# VkSemaphoreGetZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkSemaphoreGetZirconHandleInfoFUCHSIA - Structure describing a Zircon
event handle semaphore export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreGetZirconHandleInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_external_semaphore
typedef struct VkSemaphoreGetZirconHandleInfoFUCHSIA {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkSemaphoreGetZirconHandleInfoFUCHSIA;
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

The properties of the Zircon event handle returned depend on the value
of `handleType`. See
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
for a description of the properties of the defined external semaphore
handle types.

Valid Usage

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04758"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04758"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04758  
  `handleType` **must** have been included in
  [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html)::`handleTypes`
  when `semaphore`’s current payload was created

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04759"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04759"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04759  
  `semaphore` **must** not currently have its payload replaced by an
  imported payload as described below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a> unless
  that imported payload’s handle type was included in
  [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)::`exportFromImportedHandleTypes`
  for `handleType`

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04760"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04760"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04760  
  If `handleType` refers to a handle type with copy payload transference
  semantics, as defined below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a>, there
  **must** be no queue waiting on `semaphore`

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04761"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04761"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04761  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `semaphore` **must** be signaled, or have an associated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> pending
  execution

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04762"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04762"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04762  
  `handleType` **must** be defined as a Zircon event handle

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04763"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04763"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04763  
  `semaphore` **must** have been created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-sType-sType"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-sType-sType"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SEMAPHORE_GET_ZIRCON_HANDLE_INFO_FUCHSIA`

- <a href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-pNext-pNext"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-parameter"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-parameter"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a
  href="#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-parameter"
  id="VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-parameter"></a>
  VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_semaphore.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreZirconHandleFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreGetZirconHandleInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

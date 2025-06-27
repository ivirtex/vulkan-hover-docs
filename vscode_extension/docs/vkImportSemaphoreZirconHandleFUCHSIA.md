# vkImportSemaphoreZirconHandleFUCHSIA(3) Manual Page

## Name

vkImportSemaphoreZirconHandleFUCHSIA - Import a semaphore from a Zircon
event handle



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a semaphore payload from a Zircon event handle, call:

``` c
// Provided by VK_FUCHSIA_external_semaphore
VkResult vkImportSemaphoreZirconHandleFUCHSIA(
    VkDevice                                    device,
    const VkImportSemaphoreZirconHandleInfoFUCHSIA* pImportSemaphoreZirconHandleInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the semaphore.

- `pImportSemaphoreZirconHandleInfo` is a pointer to a
  [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)
  structure specifying the semaphore and import parameters.

## <a href="#_description" class="anchor"></a>Description

Importing a semaphore payload from a Zircon event handle transfers
ownership of the handle from the application to the Vulkan
implementation. The application **must** not perform any operations on
the handle after a successful import.

Applications **can** import the same semaphore payload into multiple
instances of Vulkan, into the same instance from which it was exported,
and multiple times into a given Vulkan instance.

Valid Usage

- <a href="#VUID-vkImportSemaphoreZirconHandleFUCHSIA-semaphore-04764"
  id="VUID-vkImportSemaphoreZirconHandleFUCHSIA-semaphore-04764"></a>
  VUID-vkImportSemaphoreZirconHandleFUCHSIA-semaphore-04764  
  `semaphore` **must** not be associated with any queue command that has
  not yet completed execution on that queue

Valid Usage (Implicit)

- <a href="#VUID-vkImportSemaphoreZirconHandleFUCHSIA-device-parameter"
  id="VUID-vkImportSemaphoreZirconHandleFUCHSIA-device-parameter"></a>
  VUID-vkImportSemaphoreZirconHandleFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkImportSemaphoreZirconHandleFUCHSIA-pImportSemaphoreZirconHandleInfo-parameter"
  id="VUID-vkImportSemaphoreZirconHandleFUCHSIA-pImportSemaphoreZirconHandleInfo-parameter"></a>
  VUID-vkImportSemaphoreZirconHandleFUCHSIA-pImportSemaphoreZirconHandleInfo-parameter  
  `pImportSemaphoreZirconHandleInfo` **must** be a valid pointer to a
  valid
  [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_semaphore.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkImportSemaphoreZirconHandleFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

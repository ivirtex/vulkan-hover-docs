# vkImportSemaphoreZirconHandleFUCHSIA(3) Manual Page

## Name

vkImportSemaphoreZirconHandleFUCHSIA - Import a semaphore from a Zircon event handle



## [](#_c_specification)C Specification

To import a semaphore payload from a Zircon event handle, call:

```c++
// Provided by VK_FUCHSIA_external_semaphore
VkResult vkImportSemaphoreZirconHandleFUCHSIA(
    VkDevice                                    device,
    const VkImportSemaphoreZirconHandleInfoFUCHSIA* pImportSemaphoreZirconHandleInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the semaphore.
- `pImportSemaphoreZirconHandleInfo` is a pointer to a [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html) structure specifying the semaphore and import parameters.

## [](#_description)Description

Importing a semaphore payload from a Zircon event handle transfers ownership of the handle from the application to the Vulkan implementation. The application **must** not perform any operations on the handle after a successful import.

Applications **can** import the same semaphore payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance.

Valid Usage

- [](#VUID-vkImportSemaphoreZirconHandleFUCHSIA-semaphore-04764)VUID-vkImportSemaphoreZirconHandleFUCHSIA-semaphore-04764  
  `semaphore` **must** not be associated with any queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- [](#VUID-vkImportSemaphoreZirconHandleFUCHSIA-device-parameter)VUID-vkImportSemaphoreZirconHandleFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkImportSemaphoreZirconHandleFUCHSIA-pImportSemaphoreZirconHandleInfo-parameter)VUID-vkImportSemaphoreZirconHandleFUCHSIA-pImportSemaphoreZirconHandleInfo-parameter  
  `pImportSemaphoreZirconHandleInfo` **must** be a valid pointer to a valid [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_semaphore.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkImportSemaphoreZirconHandleFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# vkGetSemaphoreZirconHandleFUCHSIA(3) Manual Page

## Name

vkGetSemaphoreZirconHandleFUCHSIA - Get a Zircon event handle for a semaphore



## [](#_c_specification)C Specification

To export a Zircon event handle representing the payload of a semaphore, call:

```c++
// Provided by VK_FUCHSIA_external_semaphore
VkResult vkGetSemaphoreZirconHandleFUCHSIA(
    VkDevice                                    device,
    const VkSemaphoreGetZirconHandleInfoFUCHSIA* pGetZirconHandleInfo,
    zx_handle_t*                                pZirconHandle);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the semaphore being exported.
- `pGetZirconHandleInfo` is a pointer to a [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html) structure containing parameters of the export operation.
- `pZirconHandle` will return the Zircon event handle representing the semaphore payload.

## [](#_description)Description

Each call to `vkGetSemaphoreZirconHandleFUCHSIA` **must** create a Zircon event handle and transfer ownership of it to the application. To avoid leaking resources, the application **must** release ownership of the Zircon event handle when it is no longer needed.

Note

Ownership can be released in many ways. For example, the application can call zx\_handle\_close() on the file descriptor, or transfer ownership back to Vulkan by using the file descriptor to import a semaphore payload.

Exporting a Zircon event handle from a semaphore **may** have side effects depending on the transference of the specified handle type, as described in [Importing Semaphore State](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing).

Valid Usage (Implicit)

- [](#VUID-vkGetSemaphoreZirconHandleFUCHSIA-device-parameter)VUID-vkGetSemaphoreZirconHandleFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSemaphoreZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter)VUID-vkGetSemaphoreZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter  
  `pGetZirconHandleInfo` **must** be a valid pointer to a valid [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html) structure
- [](#VUID-vkGetSemaphoreZirconHandleFUCHSIA-pZirconHandle-parameter)VUID-vkGetSemaphoreZirconHandleFUCHSIA-pZirconHandle-parameter  
  `pZirconHandle` **must** be a valid pointer to a `zx_handle_t` value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_semaphore.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSemaphoreZirconHandleFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
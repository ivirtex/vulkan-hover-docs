# vkGetMemoryZirconHandleFUCHSIA(3) Manual Page

## Name

vkGetMemoryZirconHandleFUCHSIA - Get a Zircon handle for an external memory object



## [](#_c_specification)C Specification

To export device memory as a Zircon handle that can be used by another instance, device, or process, retrieve the handle to the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) using the command:

```c++
// Provided by VK_FUCHSIA_external_memory
VkResult vkGetMemoryZirconHandleFUCHSIA(
    VkDevice                                    device,
    const VkMemoryGetZirconHandleInfoFUCHSIA*   pGetZirconHandleInfo,
    zx_handle_t*                                pZirconHandle);
```

## [](#_parameters)Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).
- `pGetZirconHandleInfo` is a pointer to a [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html) structure.
- `pZirconHandle` is a pointer to a `zx_handle_t` which holds the resulting Zircon handle.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryZirconHandleFUCHSIA-device-parameter)VUID-vkGetMemoryZirconHandleFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter)VUID-vkGetMemoryZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter  
  `pGetZirconHandleInfo` **must** be a valid pointer to a valid [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html) structure
- [](#VUID-vkGetMemoryZirconHandleFUCHSIA-pZirconHandle-parameter)VUID-vkGetMemoryZirconHandleFUCHSIA-pZirconHandle-parameter  
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

[VK\_FUCHSIA\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_memory.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryZirconHandleFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
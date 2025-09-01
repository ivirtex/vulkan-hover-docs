# vkGetMemoryZirconHandlePropertiesFUCHSIA(3) Manual Page

## Name

vkGetMemoryZirconHandlePropertiesFUCHSIA - Get a Zircon handle properties for an external memory object



## [](#_c_specification)C Specification

To obtain the memoryTypeIndex for the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure, call `vkGetMemoryZirconHandlePropertiesFUCHSIA`:

```c++
// Provided by VK_FUCHSIA_external_memory
VkResult vkGetMemoryZirconHandlePropertiesFUCHSIA(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    zx_handle_t                                 zirconHandle,
    VkMemoryZirconHandlePropertiesFUCHSIA*      pMemoryZirconHandleProperties);
```

## [](#_parameters)Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of `zirconHandle`
- `zirconHandle` is a `zx_handle_t` (Zircon) handle to the external resource.
- `pMemoryZirconHandleProperties` is a pointer to a [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html) structure in which the result will be stored.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-04773)VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-04773  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`
- [](#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-zirconHandle-04774)VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-zirconHandle-04774  
  `zirconHandle` **must** reference a valid VMO

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-device-parameter)VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-parameter)VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value
- [](#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-pMemoryZirconHandleProperties-parameter)VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-pMemoryZirconHandleProperties-parameter  
  `pMemoryZirconHandleProperties` **must** be a valid pointer to a [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_memory.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryZirconHandlePropertiesFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
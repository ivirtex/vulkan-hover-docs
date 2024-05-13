# vkGetMemoryZirconHandlePropertiesFUCHSIA(3) Manual Page

## Name

vkGetMemoryZirconHandlePropertiesFUCHSIA - Get a Zircon handle
properties for an external memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To obtain the memoryTypeIndex for the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure, call
`vkGetMemoryZirconHandlePropertiesFUCHSIA`:

``` c
// Provided by VK_FUCHSIA_external_memory
VkResult vkGetMemoryZirconHandlePropertiesFUCHSIA(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    zx_handle_t                                 zirconHandle,
    VkMemoryZirconHandlePropertiesFUCHSIA*      pMemoryZirconHandleProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html).

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of `zirconHandle`

- `zirconHandle` is a `zx_handle_t` (Zircon) handle to the external
  resource.

- `pMemoryZirconHandleProperties` is a pointer to a
  [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)
  structure in which the result will be stored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-04773"
  id="VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-04773"></a>
  VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-04773  
  `handleType` **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`

- <a
  href="#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-zirconHandle-04774"
  id="VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-zirconHandle-04774"></a>
  VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-zirconHandle-04774  
  `zirconHandle` must reference a valid VMO

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-device-parameter"
  id="VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-device-parameter"></a>
  VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-parameter"
  id="VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-parameter"></a>
  VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

- <a
  href="#VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-pMemoryZirconHandleProperties-parameter"
  id="VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-pMemoryZirconHandleProperties-parameter"></a>
  VUID-vkGetMemoryZirconHandlePropertiesFUCHSIA-pMemoryZirconHandleProperties-parameter  
  `pMemoryZirconHandleProperties` **must** be a valid pointer to a
  [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_memory.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryZirconHandlePropertiesFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

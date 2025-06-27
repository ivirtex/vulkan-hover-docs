# vkGetMemoryZirconHandleFUCHSIA(3) Manual Page

## Name

vkGetMemoryZirconHandleFUCHSIA - Get a Zircon handle for an external
memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export device memory as a Zircon handle that can be used by another
instance, device, or process, the handle to the
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) must be retrieved using
[vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandleFUCHSIA.html):

``` c
// Provided by VK_FUCHSIA_external_memory
VkResult vkGetMemoryZirconHandleFUCHSIA(
    VkDevice                                    device,
    const VkMemoryGetZirconHandleInfoFUCHSIA*   pGetZirconHandleInfo,
    zx_handle_t*                                pZirconHandle);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html).

- `pGetZirconHandleInfo` is a pointer to a
  [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html)
  structure.

- `pZirconHandle` is a pointer to a `zx_handle_t` which holds the
  resulting Zircon handle.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryZirconHandleFUCHSIA-device-parameter"
  id="VUID-vkGetMemoryZirconHandleFUCHSIA-device-parameter"></a>
  VUID-vkGetMemoryZirconHandleFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetMemoryZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter"
  id="VUID-vkGetMemoryZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter"></a>
  VUID-vkGetMemoryZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter  
  `pGetZirconHandleInfo` **must** be a valid pointer to a valid
  [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html)
  structure

- <a href="#VUID-vkGetMemoryZirconHandleFUCHSIA-pZirconHandle-parameter"
  id="VUID-vkGetMemoryZirconHandleFUCHSIA-pZirconHandle-parameter"></a>
  VUID-vkGetMemoryZirconHandleFUCHSIA-pZirconHandle-parameter  
  `pZirconHandle` **must** be a valid pointer to a `zx_handle_t` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_memory.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryZirconHandleFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

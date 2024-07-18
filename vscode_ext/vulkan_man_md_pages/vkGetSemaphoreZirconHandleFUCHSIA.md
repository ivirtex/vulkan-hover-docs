# vkGetSemaphoreZirconHandleFUCHSIA(3) Manual Page

## Name

vkGetSemaphoreZirconHandleFUCHSIA - Get a Zircon event handle for a
semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a Zircon event handle representing the payload of a semaphore,
call:

``` c
// Provided by VK_FUCHSIA_external_semaphore
VkResult vkGetSemaphoreZirconHandleFUCHSIA(
    VkDevice                                    device,
    const VkSemaphoreGetZirconHandleInfoFUCHSIA* pGetZirconHandleInfo,
    zx_handle_t*                                pZirconHandle);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the semaphore being
  exported.

- `pGetZirconHandleInfo` is a pointer to a
  [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)
  structure containing parameters of the export operation.

- `pZirconHandle` will return the Zircon event handle representing the
  semaphore payload.

## <a href="#_description" class="anchor"></a>Description

Each call to `vkGetSemaphoreZirconHandleFUCHSIA` **must** create a
Zircon event handle and transfer ownership of it to the application. To
avoid leaking resources, the application **must** release ownership of
the Zircon event handle when it is no longer needed.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Ownership can be released in many ways. For example, the application
can call zx_handle_close() on the file descriptor, or transfer ownership
back to Vulkan by using the file descriptor to import a semaphore
payload.</p></td>
</tr>
</tbody>
</table>

Exporting a Zircon event handle from a semaphore **may** have side
effects depending on the transference of the specified handle type, as
described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
target="_blank" rel="noopener">Importing Semaphore State</a>.

Valid Usage (Implicit)

- <a href="#VUID-vkGetSemaphoreZirconHandleFUCHSIA-device-parameter"
  id="VUID-vkGetSemaphoreZirconHandleFUCHSIA-device-parameter"></a>
  VUID-vkGetSemaphoreZirconHandleFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetSemaphoreZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter"
  id="VUID-vkGetSemaphoreZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter"></a>
  VUID-vkGetSemaphoreZirconHandleFUCHSIA-pGetZirconHandleInfo-parameter  
  `pGetZirconHandleInfo` **must** be a valid pointer to a valid
  [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)
  structure

- <a
  href="#VUID-vkGetSemaphoreZirconHandleFUCHSIA-pZirconHandle-parameter"
  id="VUID-vkGetSemaphoreZirconHandleFUCHSIA-pZirconHandle-parameter"></a>
  VUID-vkGetSemaphoreZirconHandleFUCHSIA-pZirconHandle-parameter  
  `pZirconHandle` **must** be a valid pointer to a `zx_handle_t` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_semaphore.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSemaphoreZirconHandleFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# vkGetMemoryWin32HandleKHR(3) Manual Page

## Name

vkGetMemoryWin32HandleKHR - Get a Windows HANDLE for a memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a Windows handle representing the payload of a Vulkan device
memory object, call:

``` c
// Provided by VK_KHR_external_memory_win32
VkResult vkGetMemoryWin32HandleKHR(
    VkDevice                                    device,
    const VkMemoryGetWin32HandleInfoKHR*        pGetWin32HandleInfo,
    HANDLE*                                     pHandle);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the device memory being
  exported.

- `pGetWin32HandleInfo` is a pointer to a
  [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetWin32HandleInfoKHR.html)
  structure containing parameters of the export operation.

- `pHandle` will return the Windows handle representing the payload of
  the device memory object.

## <a href="#_description" class="anchor"></a>Description

For handle types defined as NT handles, the handles returned by
`vkGetMemoryWin32HandleKHR` are owned by the application and hold a
reference to their payload. To avoid leaking resources, the application
**must** release ownership of them using the `CloseHandle` system call
when they are no longer needed.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Non-NT handle types do not add a reference to their associated
payload. If the original object owning the payload is destroyed, all
resources and handles sharing that payload will become invalid.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryWin32HandleKHR-device-parameter"
  id="VUID-vkGetMemoryWin32HandleKHR-device-parameter"></a>
  VUID-vkGetMemoryWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryWin32HandleKHR-pGetWin32HandleInfo-parameter"
  id="VUID-vkGetMemoryWin32HandleKHR-pGetWin32HandleInfo-parameter"></a>
  VUID-vkGetMemoryWin32HandleKHR-pGetWin32HandleInfo-parameter  
  `pGetWin32HandleInfo` **must** be a valid pointer to a valid
  [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetWin32HandleInfoKHR.html)
  structure

- <a href="#VUID-vkGetMemoryWin32HandleKHR-pHandle-parameter"
  id="VUID-vkGetMemoryWin32HandleKHR-pHandle-parameter"></a>
  VUID-vkGetMemoryWin32HandleKHR-pHandle-parameter  
  `pHandle` **must** be a valid pointer to a `HANDLE` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetWin32HandleInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryWin32HandleKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

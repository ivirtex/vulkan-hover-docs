# vkGetFenceFdKHR(3) Manual Page

## Name

vkGetFenceFdKHR - Get a POSIX file descriptor handle for a fence



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a POSIX file descriptor representing the payload of a fence,
call:

``` c
// Provided by VK_KHR_external_fence_fd
VkResult vkGetFenceFdKHR(
    VkDevice                                    device,
    const VkFenceGetFdInfoKHR*                  pGetFdInfo,
    int*                                        pFd);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the fence being exported.

- `pGetFdInfo` is a pointer to a
  [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetFdInfoKHR.html) structure containing
  parameters of the export operation.

- `pFd` will return the file descriptor representing the fence payload.

## <a href="#_description" class="anchor"></a>Description

Each call to `vkGetFenceFdKHR` **must** create a new file descriptor and
transfer ownership of it to the application. To avoid leaking resources,
the application **must** release ownership of the file descriptor when
it is no longer needed.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Ownership can be released in many ways. For example, the application
can call <code>close</code>() on the file descriptor, or transfer
ownership back to Vulkan by using the file descriptor to import a fence
payload.</p></td>
</tr>
</tbody>
</table>

If `pGetFdInfo->handleType` is
`VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT` and the fence is signaled at
the time `vkGetFenceFdKHR` is called, `pFd` **may** return the value
`-1` instead of a valid file descriptor.

Where supported by the operating system, the implementation **must** set
the file descriptor to be closed automatically when an `execve` system
call is made.

Exporting a file descriptor from a fence **may** have side effects
depending on the transference of the specified handle type, as described
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-importing"
target="_blank" rel="noopener">Importing Fence State</a>.

Valid Usage (Implicit)

- <a href="#VUID-vkGetFenceFdKHR-device-parameter"
  id="VUID-vkGetFenceFdKHR-device-parameter"></a>
  VUID-vkGetFenceFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetFenceFdKHR-pGetFdInfo-parameter"
  id="VUID-vkGetFenceFdKHR-pGetFdInfo-parameter"></a>
  VUID-vkGetFenceFdKHR-pGetFdInfo-parameter  
  `pGetFdInfo` **must** be a valid pointer to a valid
  [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetFdInfoKHR.html) structure

- <a href="#VUID-vkGetFenceFdKHR-pFd-parameter"
  id="VUID-vkGetFenceFdKHR-pFd-parameter"></a>
  VUID-vkGetFenceFdKHR-pFd-parameter  
  `pFd` **must** be a valid pointer to an `int` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_fd.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetFdInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetFenceFdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

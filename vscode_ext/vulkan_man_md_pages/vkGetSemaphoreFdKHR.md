# vkGetSemaphoreFdKHR(3) Manual Page

## Name

vkGetSemaphoreFdKHR - Get a POSIX file descriptor handle for a semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a POSIX file descriptor representing the payload of a
semaphore, call:

``` c
// Provided by VK_KHR_external_semaphore_fd
VkResult vkGetSemaphoreFdKHR(
    VkDevice                                    device,
    const VkSemaphoreGetFdInfoKHR*              pGetFdInfo,
    int*                                        pFd);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the semaphore being
  exported.

- `pGetFdInfo` is a pointer to a
  [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetFdInfoKHR.html) structure
  containing parameters of the export operation.

- `pFd` will return the file descriptor representing the semaphore
  payload.

## <a href="#_description" class="anchor"></a>Description

Each call to `vkGetSemaphoreFdKHR` **must** create a new file descriptor
and transfer ownership of it to the application. To avoid leaking
resources, the application **must** release ownership of the file
descriptor when it is no longer needed.

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
ownership back to Vulkan by using the file descriptor to import a
semaphore payload.</p></td>
</tr>
</tbody>
</table>

Where supported by the operating system, the implementation **must** set
the file descriptor to be closed automatically when an `execve` system
call is made.

Exporting a file descriptor from a semaphore **may** have side effects
depending on the transference of the specified handle type, as described
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
target="_blank" rel="noopener">Importing Semaphore State</a>.

Valid Usage (Implicit)

- <a href="#VUID-vkGetSemaphoreFdKHR-device-parameter"
  id="VUID-vkGetSemaphoreFdKHR-device-parameter"></a>
  VUID-vkGetSemaphoreFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetSemaphoreFdKHR-pGetFdInfo-parameter"
  id="VUID-vkGetSemaphoreFdKHR-pGetFdInfo-parameter"></a>
  VUID-vkGetSemaphoreFdKHR-pGetFdInfo-parameter  
  `pGetFdInfo` **must** be a valid pointer to a valid
  [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetFdInfoKHR.html) structure

- <a href="#VUID-vkGetSemaphoreFdKHR-pFd-parameter"
  id="VUID-vkGetSemaphoreFdKHR-pFd-parameter"></a>
  VUID-vkGetSemaphoreFdKHR-pFd-parameter  
  `pFd` **must** be a valid pointer to an `int` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_fd.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetFdInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSemaphoreFdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

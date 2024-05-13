# vkGetFenceWin32HandleKHR(3) Manual Page

## Name

vkGetFenceWin32HandleKHR - Get a Windows HANDLE for a fence



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a Windows handle representing the state of a fence, call:

``` c
// Provided by VK_KHR_external_fence_win32
VkResult vkGetFenceWin32HandleKHR(
    VkDevice                                    device,
    const VkFenceGetWin32HandleInfoKHR*         pGetWin32HandleInfo,
    HANDLE*                                     pHandle);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the fence being exported.

- `pGetWin32HandleInfo` is a pointer to a
  [VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html)
  structure containing parameters of the export operation.

- `pHandle` will return the Windows handle representing the fence state.

## <a href="#_description" class="anchor"></a>Description

For handle types defined as NT handles, the handles returned by
`vkGetFenceWin32HandleKHR` are owned by the application. To avoid
leaking resources, the application **must** release ownership of them
using the `CloseHandle` system call when they are no longer needed.

Exporting a Windows handle from a fence **may** have side effects
depending on the transference of the specified handle type, as described
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-importing"
target="_blank" rel="noopener">Importing Fence Payloads</a>.

Valid Usage (Implicit)

- <a href="#VUID-vkGetFenceWin32HandleKHR-device-parameter"
  id="VUID-vkGetFenceWin32HandleKHR-device-parameter"></a>
  VUID-vkGetFenceWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetFenceWin32HandleKHR-pGetWin32HandleInfo-parameter"
  id="VUID-vkGetFenceWin32HandleKHR-pGetWin32HandleInfo-parameter"></a>
  VUID-vkGetFenceWin32HandleKHR-pGetWin32HandleInfo-parameter  
  `pGetWin32HandleInfo` **must** be a valid pointer to a valid
  [VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html)
  structure

- <a href="#VUID-vkGetFenceWin32HandleKHR-pHandle-parameter"
  id="VUID-vkGetFenceWin32HandleKHR-pHandle-parameter"></a>
  VUID-vkGetFenceWin32HandleKHR-pHandle-parameter  
  `pHandle` **must** be a valid pointer to a `HANDLE` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_win32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetFenceWin32HandleKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

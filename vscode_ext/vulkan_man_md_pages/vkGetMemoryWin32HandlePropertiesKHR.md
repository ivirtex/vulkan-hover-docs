# vkGetMemoryWin32HandlePropertiesKHR(3) Manual Page

## Name

vkGetMemoryWin32HandlePropertiesKHR - Get Properties of External Memory
Win32 Handles



## <a href="#_c_specification" class="anchor"></a>C Specification

Windows memory handles compatible with Vulkan **may** also be created by
non-Vulkan APIs using methods beyond the scope of this specification. To
determine the correct parameters to use when importing such handles,
call:

``` c
// Provided by VK_KHR_external_memory_win32
VkResult vkGetMemoryWin32HandlePropertiesKHR(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    HANDLE                                      handle,
    VkMemoryWin32HandlePropertiesKHR*           pMemoryWin32HandleProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that will be importing `handle`.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of the handle `handle`.

- `handle` is the handle which will be imported.

- `pMemoryWin32HandleProperties` is a pointer to a
  [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryWin32HandlePropertiesKHR.html)
  structure in which properties of `handle` are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetMemoryWin32HandlePropertiesKHR-handle-00665"
  id="VUID-vkGetMemoryWin32HandlePropertiesKHR-handle-00665"></a>
  VUID-vkGetMemoryWin32HandlePropertiesKHR-handle-00665  
  `handle` **must** point to a valid Windows memory handle

- <a href="#VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-00666"
  id="VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-00666"></a>
  VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-00666  
  `handleType` **must** not be one of the handle types defined as opaque

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryWin32HandlePropertiesKHR-device-parameter"
  id="VUID-vkGetMemoryWin32HandlePropertiesKHR-device-parameter"></a>
  VUID-vkGetMemoryWin32HandlePropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-parameter"
  id="VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-parameter"></a>
  VUID-vkGetMemoryWin32HandlePropertiesKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

- <a
  href="#VUID-vkGetMemoryWin32HandlePropertiesKHR-pMemoryWin32HandleProperties-parameter"
  id="VUID-vkGetMemoryWin32HandlePropertiesKHR-pMemoryWin32HandleProperties-parameter"></a>
  VUID-vkGetMemoryWin32HandlePropertiesKHR-pMemoryWin32HandleProperties-parameter  
  `pMemoryWin32HandleProperties` **must** be a valid pointer to a
  [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryWin32HandlePropertiesKHR.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryWin32HandlePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryWin32HandlePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

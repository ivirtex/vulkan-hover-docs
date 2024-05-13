# vkGetMemoryFdPropertiesKHR(3) Manual Page

## Name

vkGetMemoryFdPropertiesKHR - Get Properties of External Memory File
Descriptors



## <a href="#_c_specification" class="anchor"></a>C Specification

POSIX file descriptor memory handles compatible with Vulkan **may** also
be created by non-Vulkan APIs using methods beyond the scope of this
specification. To determine the correct parameters to use when importing
such handles, call:

``` c
// Provided by VK_KHR_external_memory_fd
VkResult vkGetMemoryFdPropertiesKHR(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    int                                         fd,
    VkMemoryFdPropertiesKHR*                    pMemoryFdProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that will be importing `fd`.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of the handle `fd`.

- `fd` is the handle which will be imported.

- `pMemoryFdProperties` is a pointer to a
  [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryFdPropertiesKHR.html) structure in
  which the properties of the handle `fd` are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetMemoryFdPropertiesKHR-fd-00673"
  id="VUID-vkGetMemoryFdPropertiesKHR-fd-00673"></a>
  VUID-vkGetMemoryFdPropertiesKHR-fd-00673  
  `fd` **must** point to a valid POSIX file descriptor memory handle

- <a href="#VUID-vkGetMemoryFdPropertiesKHR-handleType-00674"
  id="VUID-vkGetMemoryFdPropertiesKHR-handleType-00674"></a>
  VUID-vkGetMemoryFdPropertiesKHR-handleType-00674  
  `handleType` **must** not be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryFdPropertiesKHR-device-parameter"
  id="VUID-vkGetMemoryFdPropertiesKHR-device-parameter"></a>
  VUID-vkGetMemoryFdPropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryFdPropertiesKHR-handleType-parameter"
  id="VUID-vkGetMemoryFdPropertiesKHR-handleType-parameter"></a>
  VUID-vkGetMemoryFdPropertiesKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

- <a href="#VUID-vkGetMemoryFdPropertiesKHR-pMemoryFdProperties-parameter"
  id="VUID-vkGetMemoryFdPropertiesKHR-pMemoryFdProperties-parameter"></a>
  VUID-vkGetMemoryFdPropertiesKHR-pMemoryFdProperties-parameter  
  `pMemoryFdProperties` **must** be a valid pointer to a
  [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryFdPropertiesKHR.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryFdPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryFdPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

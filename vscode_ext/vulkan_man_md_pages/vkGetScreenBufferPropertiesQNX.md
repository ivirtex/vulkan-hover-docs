# vkGetScreenBufferPropertiesQNX(3) Manual Page

## Name

vkGetScreenBufferPropertiesQNX - Get Properties of External Memory QNX
Screen Buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory parameters to use when importing a QNX Screen
buffer, call:

``` c
// Provided by VK_QNX_external_memory_screen_buffer
VkResult vkGetScreenBufferPropertiesQNX(
    VkDevice                                    device,
    const struct _screen_buffer*                buffer,
    VkScreenBufferPropertiesQNX*                pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that will be importing `buffer`.

- `buffer` is the QNX Screen buffer which will be imported.

- `pProperties` is a pointer to a
  [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferPropertiesQNX.html)
  structure in which the properties of `buffer` are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetScreenBufferPropertiesQNX-buffer-08968"
  id="VUID-vkGetScreenBufferPropertiesQNX-buffer-08968"></a>
  VUID-vkGetScreenBufferPropertiesQNX-buffer-08968  
  `buffer` **must** be a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-validity"
  target="_blank" rel="noopener">valid QNX Screen buffer</a>

Valid Usage (Implicit)

- <a href="#VUID-vkGetScreenBufferPropertiesQNX-device-parameter"
  id="VUID-vkGetScreenBufferPropertiesQNX-device-parameter"></a>
  VUID-vkGetScreenBufferPropertiesQNX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetScreenBufferPropertiesQNX-buffer-parameter"
  id="VUID-vkGetScreenBufferPropertiesQNX-buffer-parameter"></a>
  VUID-vkGetScreenBufferPropertiesQNX-buffer-parameter  
  `buffer` **must** be a valid pointer to a valid `_screen_buffer` value

- <a href="#VUID-vkGetScreenBufferPropertiesQNX-pProperties-parameter"
  id="VUID-vkGetScreenBufferPropertiesQNX-pProperties-parameter"></a>
  VUID-vkGetScreenBufferPropertiesQNX-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferPropertiesQNX.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_external_memory_screen_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_external_memory_screen_buffer.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferPropertiesQNX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetScreenBufferPropertiesQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

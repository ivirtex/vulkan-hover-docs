# vkGetMemoryHostPointerPropertiesEXT(3) Manual Page

## Name

vkGetMemoryHostPointerPropertiesEXT - Get properties of external memory
host pointer



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the correct parameters to use when importing host pointers,
call:

``` c
// Provided by VK_EXT_external_memory_host
VkResult vkGetMemoryHostPointerPropertiesEXT(
    VkDevice                                    device,
    VkExternalMemoryHandleTypeFlagBits          handleType,
    const void*                                 pHostPointer,
    VkMemoryHostPointerPropertiesEXT*           pMemoryHostPointerProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that will be importing `pHostPointer`.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of the handle `pHostPointer`.

- `pHostPointer` is the host pointer to import from.

- `pMemoryHostPointerProperties` is a pointer to a
  [VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHostPointerPropertiesEXT.html)
  structure in which the host pointer properties are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01752"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01752"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01752  
  `handleType` **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`

- <a href="#VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-01753"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-01753"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-01753  
  `pHostPointer` **must** be a pointer aligned to an integer multiple of
  `VkPhysicalDeviceExternalMemoryHostPropertiesEXT`::`minImportedHostPointerAlignment`

- <a href="#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01754"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01754"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01754  
  If `handleType` is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT`,
  `pHostPointer` **must** be a pointer to host memory

- <a href="#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01755"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01755"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-01755  
  If `handleType` is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`,
  `pHostPointer` **must** be a pointer to host mapped foreign memory

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryHostPointerPropertiesEXT-device-parameter"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-device-parameter"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-parameter"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-parameter"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

- <a
  href="#VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-parameter"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-parameter"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-pHostPointer-parameter  
  `pHostPointer` **must** be a pointer value

- <a
  href="#VUID-vkGetMemoryHostPointerPropertiesEXT-pMemoryHostPointerProperties-parameter"
  id="VUID-vkGetMemoryHostPointerPropertiesEXT-pMemoryHostPointerProperties-parameter"></a>
  VUID-vkGetMemoryHostPointerPropertiesEXT-pMemoryHostPointerProperties-parameter  
  `pMemoryHostPointerProperties` **must** be a valid pointer to a
  [VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHostPointerPropertiesEXT.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_external_memory_host](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_external_memory_host.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHostPointerPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryHostPointerPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

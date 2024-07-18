# vkGetMemoryWin32HandleNV(3) Manual Page

## Name

vkGetMemoryWin32HandleNV - Retrieve Win32 handle to a device memory
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve the handle corresponding to a device memory object created
with
[VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html)::`handleTypes`
set to include `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV` or
`VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV`, call:

``` c
// Provided by VK_NV_external_memory_win32
VkResult vkGetMemoryWin32HandleNV(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    VkExternalMemoryHandleTypeFlagsNV           handleType,
    HANDLE*                                     pHandle);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object.

- `handleType` is a bitmask of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  containing a single bit specifying the type of handle requested.

- `handle` is a pointer to a Windows `HANDLE` in which the handle is
  returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetMemoryWin32HandleNV-handleType-01326"
  id="VUID-vkGetMemoryWin32HandleNV-handleType-01326"></a>
  VUID-vkGetMemoryWin32HandleNV-handleType-01326  
  `handleType` **must** be a flag specified in
  [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html)::`handleTypes`
  when allocating `memory`

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryWin32HandleNV-device-parameter"
  id="VUID-vkGetMemoryWin32HandleNV-device-parameter"></a>
  VUID-vkGetMemoryWin32HandleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryWin32HandleNV-memory-parameter"
  id="VUID-vkGetMemoryWin32HandleNV-memory-parameter"></a>
  VUID-vkGetMemoryWin32HandleNV-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-vkGetMemoryWin32HandleNV-handleType-parameter"
  id="VUID-vkGetMemoryWin32HandleNV-handleType-parameter"></a>
  VUID-vkGetMemoryWin32HandleNV-handleType-parameter  
  `handleType` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  values

- <a href="#VUID-vkGetMemoryWin32HandleNV-handleType-requiredbitmask"
  id="VUID-vkGetMemoryWin32HandleNV-handleType-requiredbitmask"></a>
  VUID-vkGetMemoryWin32HandleNV-handleType-requiredbitmask  
  `handleType` **must** not be `0`

- <a href="#VUID-vkGetMemoryWin32HandleNV-pHandle-parameter"
  id="VUID-vkGetMemoryWin32HandleNV-pHandle-parameter"></a>
  VUID-vkGetMemoryWin32HandleNV-pHandle-parameter  
  `pHandle` **must** be a valid pointer to a `HANDLE` value

- <a href="#VUID-vkGetMemoryWin32HandleNV-memory-parent"
  id="VUID-vkGetMemoryWin32HandleNV-memory-parent"></a>
  VUID-vkGetMemoryWin32HandleNV-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_win32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryWin32HandleNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

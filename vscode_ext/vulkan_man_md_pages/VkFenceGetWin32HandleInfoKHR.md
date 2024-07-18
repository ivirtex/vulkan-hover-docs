# VkFenceGetWin32HandleInfoKHR(3) Manual Page

## Name

VkFenceGetWin32HandleInfoKHR - Structure describing a Win32 handle fence
export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFenceGetWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_fence_win32
typedef struct VkFenceGetWin32HandleInfoKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    VkFence                              fence;
    VkExternalFenceHandleTypeFlagBits    handleType;
} VkFenceGetWin32HandleInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fence` is the fence from which state will be exported.

- `handleType` is a
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value specifying the type of handle requested.

## <a href="#_description" class="anchor"></a>Description

The properties of the handle returned depend on the value of
`handleType`. See
[VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
for a description of the properties of the defined external fence handle
types.

Valid Usage

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01448"
  id="VUID-VkFenceGetWin32HandleInfoKHR-handleType-01448"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-handleType-01448  
  `handleType` **must** have been included in
  [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html)::`handleTypes`
  when the `fence`’s current payload was created

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01449"
  id="VUID-VkFenceGetWin32HandleInfoKHR-handleType-01449"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-handleType-01449  
  If `handleType` is defined as an NT handle,
  [vkGetFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceWin32HandleKHR.html) **must** be
  called no more than once for each valid unique combination of `fence`
  and `handleType`

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-fence-01450"
  id="VUID-VkFenceGetWin32HandleInfoKHR-fence-01450"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-fence-01450  
  `fence` **must** not currently have its payload replaced by an
  imported payload as described below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-importing"
  target="_blank" rel="noopener">Importing Fence Payloads</a> unless
  that imported payload’s handle type was included in
  [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)::`exportFromImportedHandleTypes`
  for `handleType`

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01451"
  id="VUID-VkFenceGetWin32HandleInfoKHR-handleType-01451"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-handleType-01451  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `fence` **must** be signaled, or have an associated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-signaling"
  target="_blank" rel="noopener">fence signal operation</a> pending
  execution

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01452"
  id="VUID-VkFenceGetWin32HandleInfoKHR-handleType-01452"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-handleType-01452  
  `handleType` **must** be defined as an NT handle or a global share
  handle

Valid Usage (Implicit)

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-sType-sType"
  id="VUID-VkFenceGetWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-pNext-pNext"
  id="VUID-VkFenceGetWin32HandleInfoKHR-pNext-pNext"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-fence-parameter"
  id="VUID-VkFenceGetWin32HandleInfoKHR-fence-parameter"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-VkFenceGetWin32HandleInfoKHR-handleType-parameter"
  id="VUID-VkFenceGetWin32HandleInfoKHR-handleType-parameter"></a>
  VUID-VkFenceGetWin32HandleInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_win32.html),
[VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFenceWin32HandleKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFenceGetWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# vkResetFences(3) Manual Page

## Name

vkResetFences - Resets one or more fence objects



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the state of fences to unsignaled from the host, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkResetFences(
    VkDevice                                    device,
    uint32_t                                    fenceCount,
    const VkFence*                              pFences);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the fences.

- `fenceCount` is the number of fences to reset.

- `pFences` is a pointer to an array of fence handles to reset.

## <a href="#_description" class="anchor"></a>Description

If any member of `pFences` currently has its <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-importing"
target="_blank" rel="noopener">payload imported</a> with temporary
permanence, that fence’s prior permanent payload is first restored. The
remaining operations described therefore operate on the restored
payload.

When [vkResetFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetFences.html) is executed on the host, it
defines a *fence unsignal operation* for each fence, which resets the
fence to the unsignaled state.

If any member of `pFences` is already in the unsignaled state when
[vkResetFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetFences.html) is executed, then
[vkResetFences](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetFences.html) has no effect on that fence.

Valid Usage

- <a href="#VUID-vkResetFences-pFences-01123"
  id="VUID-vkResetFences-pFences-01123"></a>
  VUID-vkResetFences-pFences-01123  
  Each element of `pFences` **must** not be currently associated with
  any queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- <a href="#VUID-vkResetFences-device-parameter"
  id="VUID-vkResetFences-device-parameter"></a>
  VUID-vkResetFences-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkResetFences-pFences-parameter"
  id="VUID-vkResetFences-pFences-parameter"></a>
  VUID-vkResetFences-pFences-parameter  
  `pFences` **must** be a valid pointer to an array of `fenceCount`
  valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handles

- <a href="#VUID-vkResetFences-fenceCount-arraylength"
  id="VUID-vkResetFences-fenceCount-arraylength"></a>
  VUID-vkResetFences-fenceCount-arraylength  
  `fenceCount` **must** be greater than `0`

- <a href="#VUID-vkResetFences-pFences-parent"
  id="VUID-vkResetFences-pFences-parent"></a>
  VUID-vkResetFences-pFences-parent  
  Each element of `pFences` **must** have been created, allocated, or
  retrieved from `device`

Host Synchronization

- Host access to each member of `pFences` **must** be externally
  synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkResetFences"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

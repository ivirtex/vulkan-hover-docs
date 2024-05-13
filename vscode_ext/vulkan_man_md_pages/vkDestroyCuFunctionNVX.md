# vkDestroyCuFunctionNVX(3) Manual Page

## Name

vkDestroyCuFunctionNVX - Stub description of vkDestroyCuFunctionNVX



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this command.
This section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_NVX_binary_import
void vkDestroyCuFunctionNVX(
    VkDevice                                    device,
    VkCuFunctionNVX                             function,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyCuFunctionNVX-device-parameter"
  id="VUID-vkDestroyCuFunctionNVX-device-parameter"></a>
  VUID-vkDestroyCuFunctionNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyCuFunctionNVX-function-parameter"
  id="VUID-vkDestroyCuFunctionNVX-function-parameter"></a>
  VUID-vkDestroyCuFunctionNVX-function-parameter  
  `function` **must** be a valid [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuFunctionNVX.html)
  handle

- <a href="#VUID-vkDestroyCuFunctionNVX-pAllocator-parameter"
  id="VUID-vkDestroyCuFunctionNVX-pAllocator-parameter"></a>
  VUID-vkDestroyCuFunctionNVX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyCuFunctionNVX-function-parent"
  id="VUID-vkDestroyCuFunctionNVX-function-parent"></a>
  VUID-vkDestroyCuFunctionNVX-function-parent  
  `function` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_binary_import](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_binary_import.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuFunctionNVX.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyCuFunctionNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# vkCreateCuModuleNVX(3) Manual Page

## Name

vkCreateCuModuleNVX - Stub description of vkCreateCuModuleNVX



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this command.
This section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_NVX_binary_import
VkResult vkCreateCuModuleNVX(
    VkDevice                                    device,
    const VkCuModuleCreateInfoNVX*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCuModuleNVX*                              pModule);
```

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateCuModuleNVX-device-parameter"
  id="VUID-vkCreateCuModuleNVX-device-parameter"></a>
  VUID-vkCreateCuModuleNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateCuModuleNVX-pCreateInfo-parameter"
  id="VUID-vkCreateCuModuleNVX-pCreateInfo-parameter"></a>
  VUID-vkCreateCuModuleNVX-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleCreateInfoNVX.html) structure

- <a href="#VUID-vkCreateCuModuleNVX-pAllocator-parameter"
  id="VUID-vkCreateCuModuleNVX-pAllocator-parameter"></a>
  VUID-vkCreateCuModuleNVX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateCuModuleNVX-pModule-parameter"
  id="VUID-vkCreateCuModuleNVX-pModule-parameter"></a>
  VUID-vkCreateCuModuleNVX-pModule-parameter  
  `pModule` **must** be a valid pointer to a
  [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleNVX.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_binary_import](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_binary_import.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleCreateInfoNVX.html),
[VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleNVX.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateCuModuleNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

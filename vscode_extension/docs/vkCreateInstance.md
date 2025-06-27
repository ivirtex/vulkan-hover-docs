# vkCreateInstance(3) Manual Page

## Name

vkCreateInstance - Create a new Vulkan instance



## <a href="#_c_specification" class="anchor"></a>C Specification

To create an instance object, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateInstance(
    const VkInstanceCreateInfo*                 pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkInstance*                                 pInstance);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pCreateInfo` is a pointer to a
  [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure
  controlling creation of the instance.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pInstance` points a [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle in which the
  resulting instance is returned.

## <a href="#_description" class="anchor"></a>Description

`vkCreateInstance` verifies that the requested layers exist. If not,
`vkCreateInstance` will return `VK_ERROR_LAYER_NOT_PRESENT`. Next
`vkCreateInstance` verifies that the requested extensions are supported
(e.g. in the implementation or in any enabled instance layer) and if any
requested extension is not supported, `vkCreateInstance` **must** return
`VK_ERROR_EXTENSION_NOT_PRESENT`. After verifying and enabling the
instance layers and extensions the `VkInstance` object is created and
returned to the application. If a requested extension is only supported
by a layer, both the layer and the extension need to be specified at
`vkCreateInstance` time for the creation to succeed.

Valid Usage

- <a href="#VUID-vkCreateInstance-ppEnabledExtensionNames-01388"
  id="VUID-vkCreateInstance-ppEnabledExtensionNames-01388"></a>
  VUID-vkCreateInstance-ppEnabledExtensionNames-01388  
  All <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-extensions-extensiondependencies"
  target="_blank" rel="noopener">required extensions</a> for each
  extension in the
  [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html)::`ppEnabledExtensionNames`
  list **must** also be present in that list

Valid Usage (Implicit)

- <a href="#VUID-vkCreateInstance-pCreateInfo-parameter"
  id="VUID-vkCreateInstance-pCreateInfo-parameter"></a>
  VUID-vkCreateInstance-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure

- <a href="#VUID-vkCreateInstance-pAllocator-parameter"
  id="VUID-vkCreateInstance-pAllocator-parameter"></a>
  VUID-vkCreateInstance-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateInstance-pInstance-parameter"
  id="VUID-vkCreateInstance-pInstance-parameter"></a>
  VUID-vkCreateInstance-pInstance-parameter  
  `pInstance` **must** be a valid pointer to a
  [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_LAYER_NOT_PRESENT`

- `VK_ERROR_EXTENSION_NOT_PRESENT`

- `VK_ERROR_INCOMPATIBLE_DRIVER`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html),
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateInstance"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

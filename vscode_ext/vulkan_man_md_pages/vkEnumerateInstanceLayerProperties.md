# vkEnumerateInstanceLayerProperties(3) Manual Page

## Name

vkEnumerateInstanceLayerProperties - Returns up to requested number of
global layer properties



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the available layers, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkEnumerateInstanceLayerProperties(
    uint32_t*                                   pPropertyCount,
    VkLayerProperties*                          pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pPropertyCount` is a pointer to an integer related to the number of
  layer properties available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html) structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of layer properties
available is returned in `pPropertyCount`. Otherwise, `pPropertyCount`
**must** point to a variable set by the user to the number of elements
in the `pProperties` array, and on return the variable is overwritten
with the number of structures actually written to `pProperties`. If
`pPropertyCount` is less than the number of layer properties available,
at most `pPropertyCount` structures will be written, and `VK_INCOMPLETE`
will be returned instead of `VK_SUCCESS`, to indicate that not all the
available properties were returned.

The list of available layers may change at any time due to actions
outside of the Vulkan implementation, so two calls to
`vkEnumerateInstanceLayerProperties` with the same parameters **may**
return different results, or retrieve different `pPropertyCount` values
or `pProperties` contents. Once an instance has been created, the layers
enabled for that instance will continue to be enabled and valid for the
lifetime of that instance, even if some of them become unavailable for
future instances.

Valid Usage (Implicit)

- <a
  href="#VUID-vkEnumerateInstanceLayerProperties-pPropertyCount-parameter"
  id="VUID-vkEnumerateInstanceLayerProperties-pPropertyCount-parameter"></a>
  VUID-vkEnumerateInstanceLayerProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkEnumerateInstanceLayerProperties-pProperties-parameter"
  id="VUID-vkEnumerateInstanceLayerProperties-pProperties-parameter"></a>
  VUID-vkEnumerateInstanceLayerProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html) structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumerateInstanceLayerProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# vkEnumerateInstanceExtensionProperties(3) Manual Page

## Name

vkEnumerateInstanceExtensionProperties - Returns up to requested number
of global extension properties



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the available instance extensions, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkEnumerateInstanceExtensionProperties(
    const char*                                 pLayerName,
    uint32_t*                                   pPropertyCount,
    VkExtensionProperties*                      pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pLayerName` is either `NULL` or a pointer to a null-terminated UTF-8
  string naming the layer to retrieve extensions from.

- `pPropertyCount` is a pointer to an integer related to the number of
  extension properties available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structures.

## <a href="#_description" class="anchor"></a>Description

When `pLayerName` parameter is `NULL`, only extensions provided by the
Vulkan implementation or by implicitly enabled layers are returned. When
`pLayerName` is the name of a layer, the instance extensions provided by
that layer are returned.

If `pProperties` is `NULL`, then the number of extensions properties
available is returned in `pPropertyCount`. Otherwise, `pPropertyCount`
**must** point to a variable set by the application to the number of
elements in the `pProperties` array, and on return the variable is
overwritten with the number of structures actually written to
`pProperties`. If `pPropertyCount` is less than the number of extension
properties available, at most `pPropertyCount` structures will be
written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`,
to indicate that not all the available properties were returned.

Because the list of available layers may change externally between calls
to
[vkEnumerateInstanceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html),
two calls may retrieve different results if a `pLayerName` is available
in one call but not in another. The extensions supported by a layer may
also change between two calls, e.g. if the layer implementation is
replaced by a different version between those calls.

Implementations **must** not advertise any pair of extensions that
cannot be enabled together due to behavioral differences, or any
extension that cannot be enabled against the advertised version.

Valid Usage (Implicit)

- <a
  href="#VUID-vkEnumerateInstanceExtensionProperties-pLayerName-parameter"
  id="VUID-vkEnumerateInstanceExtensionProperties-pLayerName-parameter"></a>
  VUID-vkEnumerateInstanceExtensionProperties-pLayerName-parameter  
  If `pLayerName` is not `NULL`, `pLayerName` **must** be a
  null-terminated UTF-8 string

- <a
  href="#VUID-vkEnumerateInstanceExtensionProperties-pPropertyCount-parameter"
  id="VUID-vkEnumerateInstanceExtensionProperties-pPropertyCount-parameter"></a>
  VUID-vkEnumerateInstanceExtensionProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkEnumerateInstanceExtensionProperties-pProperties-parameter"
  id="VUID-vkEnumerateInstanceExtensionProperties-pProperties-parameter"></a>
  VUID-vkEnumerateInstanceExtensionProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_LAYER_NOT_PRESENT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumerateInstanceExtensionProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

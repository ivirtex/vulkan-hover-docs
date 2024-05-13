# VkLayerProperties(3) Manual Page

## Name

VkLayerProperties - Structure specifying layer properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkLayerProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkLayerProperties {
    char        layerName[VK_MAX_EXTENSION_NAME_SIZE];
    uint32_t    specVersion;
    uint32_t    implementationVersion;
    char        description[VK_MAX_DESCRIPTION_SIZE];
} VkLayerProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `layerName` is an array of `VK_MAX_EXTENSION_NAME_SIZE` `char`
  containing a null-terminated UTF-8 string which is the name of the
  layer. Use this name in the `ppEnabledLayerNames` array passed in the
  [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure to enable
  this layer for an instance.

- `specVersion` is the Vulkan version the layer was written to, encoded
  as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers</a>.

- `implementationVersion` is the version of this layer. It is an
  integer, increasing with backward compatible changes.

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which provides additional
  details that **can** be used by the application to identify the layer.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkEnumerateDeviceLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceLayerProperties.html),
[vkEnumerateInstanceLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceLayerProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLayerProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# PFN_vkDebugUtilsMessengerCallbackEXT(3) Manual Page

## Name

PFN_vkDebugUtilsMessengerCallbackEXT - Application-defined debug
messenger callback function



## <a href="#_c_specification" class="anchor"></a>C Specification

The prototype for the
[VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)::`pfnUserCallback`
function implemented by the application is:

``` c
// Provided by VK_EXT_debug_utils
typedef VkBool32 (VKAPI_PTR *PFN_vkDebugUtilsMessengerCallbackEXT)(
    VkDebugUtilsMessageSeverityFlagBitsEXT           messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT                  messageTypes,
    const VkDebugUtilsMessengerCallbackDataEXT*      pCallbackData,
    void*                                            pUserData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `messageSeverity` specifies the
  [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html)
  that triggered this callback.

- `messageTypes` is a bitmask of
  [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html)
  specifying which type of event(s) triggered this callback.

- `pCallbackData` contains all the callback related data in the
  [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html)
  structure.

- `pUserData` is the user data provided when the
  [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) was created.

## <a href="#_description" class="anchor"></a>Description

The callback returns a [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), which is interpreted
in a layer-specified manner. The application **should** always return
`VK_FALSE`. The `VK_TRUE` value is reserved for use in layer
development.

Valid Usage

- <a href="#VUID-PFN_vkDebugUtilsMessengerCallbackEXT-None-04769"
  id="VUID-PFN_vkDebugUtilsMessengerCallbackEXT-None-04769"></a>
  VUID-PFN_vkDebugUtilsMessengerCallbackEXT-None-04769  
  The callback **must** not make calls to any Vulkan commands

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html),
[VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagsEXT.html),
[VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html),
[VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PFN_vkDebugUtilsMessengerCallbackEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

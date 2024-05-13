# VkDebugUtilsMessengerCreateInfoEXT(3) Manual Page

## Name

VkDebugUtilsMessengerCreateInfoEXT - Structure specifying parameters of
a newly created debug messenger



## <a href="#_c_specification" class="anchor"></a>C Specification

The definition of `VkDebugUtilsMessengerCreateInfoEXT` is:

``` c
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsMessengerCreateInfoEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    VkDebugUtilsMessengerCreateFlagsEXT     flags;
    VkDebugUtilsMessageSeverityFlagsEXT     messageSeverity;
    VkDebugUtilsMessageTypeFlagsEXT         messageType;
    PFN_vkDebugUtilsMessengerCallbackEXT    pfnUserCallback;
    void*                                   pUserData;
} VkDebugUtilsMessengerCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is `0` and is reserved for future use.

- `messageSeverity` is a bitmask of
  [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html)
  specifying which severity of event(s) will cause this callback to be
  called.

- `messageType` is a bitmask of
  [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html)
  specifying which type of event(s) will cause this callback to be
  called.

- `pfnUserCallback` is the application callback function to call.

- `pUserData` is user data to be passed to the callback.

## <a href="#_description" class="anchor"></a>Description

For each `VkDebugUtilsMessengerEXT` that is created the
`VkDebugUtilsMessengerCreateInfoEXT`::`messageSeverity` and
`VkDebugUtilsMessengerCreateInfoEXT`::`messageType` determine when that
`VkDebugUtilsMessengerCreateInfoEXT`::`pfnUserCallback` is called. The
process to determine if the user’s `pfnUserCallback` is triggered when
an event occurs is as follows:

1.  The implementation will perform a bitwise AND of the event’s
    [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html)
    with the `messageSeverity` provided during creation of the
    [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) object.

    1.  If the value is 0, the message is skipped.

2.  The implementation will perform bitwise AND of the event’s
    [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html)
    with the `messageType` provided during the creation of the
    [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) object.

    1.  If the value is 0, the message is skipped.

3.  The callback will trigger a debug message for the current event

The callback will come directly from the component that detected the
event, unless some other layer intercepts the calls for its own purposes
(filter them in a different way, log to a system error log, etc.).

An application **can** receive multiple callbacks if multiple
`VkDebugUtilsMessengerEXT` objects are created. A callback will always
be executed in the same thread as the originating Vulkan call.

A callback **can** be called from multiple threads simultaneously (if
the application is making Vulkan calls from multiple threads).

Valid Usage (Implicit)

- <a href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-sType-sType"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-sType-sType"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT`

- <a href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-parameter"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-parameter"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-parameter  
  `messageSeverity` **must** be a valid combination of
  [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html)
  values

- <a
  href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-requiredbitmask"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-requiredbitmask"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-requiredbitmask  
  `messageSeverity` **must** not be `0`

- <a href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-parameter"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-parameter"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-parameter  
  `messageType` **must** be a valid combination of
  [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html)
  values

- <a
  href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-requiredbitmask"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-requiredbitmask"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-requiredbitmask  
  `messageType` **must** not be `0`

- <a
  href="#VUID-VkDebugUtilsMessengerCreateInfoEXT-pfnUserCallback-parameter"
  id="VUID-VkDebugUtilsMessengerCreateInfoEXT-pfnUserCallback-parameter"></a>
  VUID-VkDebugUtilsMessengerCreateInfoEXT-pfnUserCallback-parameter  
  `pfnUserCallback` **must** be a valid
  [PFN_vkDebugUtilsMessengerCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDebugUtilsMessengerCallbackEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkDebugUtilsMessengerCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDebugUtilsMessengerCallbackEXT.html),
[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsMessageSeverityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagsEXT.html),
[VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagsEXT.html),
[VkDebugUtilsMessengerCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsMessengerCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

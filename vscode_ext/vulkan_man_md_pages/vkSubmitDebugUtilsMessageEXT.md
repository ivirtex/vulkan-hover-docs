# vkSubmitDebugUtilsMessageEXT(3) Manual Page

## Name

vkSubmitDebugUtilsMessageEXT - Inject a message into a debug stream



## <a href="#_c_specification" class="anchor"></a>C Specification

To intentionally submit a debug message, call:

``` c
// Provided by VK_EXT_debug_utils
void vkSubmitDebugUtilsMessageEXT(
    VkInstance                                  instance,
    VkDebugUtilsMessageSeverityFlagBitsEXT      messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT             messageTypes,
    const VkDebugUtilsMessengerCallbackDataEXT* pCallbackData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the debug stream’s [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html).

- `messageSeverity` is a
  [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html)
  value specifying the severity of this event/message.

- `messageTypes` is a bitmask of
  [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html)
  specifying which type of event(s) to identify with this message.

- `pCallbackData` contains all the callback related data in the
  [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html)
  structure.

## <a href="#_description" class="anchor"></a>Description

The call will propagate through the layers and generate callback(s) as
indicated by the message’s flags. The parameters are passed on to the
callback in addition to the `pUserData` value that was defined at the
time the messenger was registered.

Valid Usage

- <a href="#VUID-vkSubmitDebugUtilsMessageEXT-objectType-02591"
  id="VUID-vkSubmitDebugUtilsMessageEXT-objectType-02591"></a>
  VUID-vkSubmitDebugUtilsMessageEXT-objectType-02591  
  The `objectType` member of each element of `pCallbackData->pObjects`
  **must** not be `VK_OBJECT_TYPE_UNKNOWN`

Valid Usage (Implicit)

- <a href="#VUID-vkSubmitDebugUtilsMessageEXT-instance-parameter"
  id="VUID-vkSubmitDebugUtilsMessageEXT-instance-parameter"></a>
  VUID-vkSubmitDebugUtilsMessageEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkSubmitDebugUtilsMessageEXT-messageSeverity-parameter"
  id="VUID-vkSubmitDebugUtilsMessageEXT-messageSeverity-parameter"></a>
  VUID-vkSubmitDebugUtilsMessageEXT-messageSeverity-parameter  
  `messageSeverity` **must** be a valid
  [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html)
  value

- <a href="#VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-parameter"
  id="VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-parameter"></a>
  VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-parameter  
  `messageTypes` **must** be a valid combination of
  [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html)
  values

- <a
  href="#VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-requiredbitmask"
  id="VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-requiredbitmask"></a>
  VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-requiredbitmask  
  `messageTypes` **must** not be `0`

- <a href="#VUID-vkSubmitDebugUtilsMessageEXT-pCallbackData-parameter"
  id="VUID-vkSubmitDebugUtilsMessageEXT-pCallbackData-parameter"></a>
  VUID-vkSubmitDebugUtilsMessageEXT-pCallbackData-parameter  
  `pCallbackData` **must** be a valid pointer to a valid
  [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html),
[VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagsEXT.html),
[VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSubmitDebugUtilsMessageEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

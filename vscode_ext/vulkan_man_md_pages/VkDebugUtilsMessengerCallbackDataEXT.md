# VkDebugUtilsMessengerCallbackDataEXT(3) Manual Page

## Name

VkDebugUtilsMessengerCallbackDataEXT - Structure specifying parameters
returned to the callback



## <a href="#_c_specification" class="anchor"></a>C Specification

The definition of `VkDebugUtilsMessengerCallbackDataEXT` is:

``` c
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsMessengerCallbackDataEXT {
    VkStructureType                              sType;
    const void*                                  pNext;
    VkDebugUtilsMessengerCallbackDataFlagsEXT    flags;
    const char*                                  pMessageIdName;
    int32_t                                      messageIdNumber;
    const char*                                  pMessage;
    uint32_t                                     queueLabelCount;
    const VkDebugUtilsLabelEXT*                  pQueueLabels;
    uint32_t                                     cmdBufLabelCount;
    const VkDebugUtilsLabelEXT*                  pCmdBufLabels;
    uint32_t                                     objectCount;
    const VkDebugUtilsObjectNameInfoEXT*         pObjects;
} VkDebugUtilsMessengerCallbackDataEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is `0` and is reserved for future use.

- `pMessageIdName` is `NULL` or a null-terminated UTF-8 string that
  identifies the particular message ID that is associated with the
  provided message. If the message corresponds to a validation layer
  message, then this string may contain the portion of the Vulkan
  specification that is believed to have been violated.

- `messageIdNumber` is the ID number of the triggering message. If the
  message corresponds to a validation layer message, then this number is
  related to the internal number associated with the message being
  triggered.

- `pMessage` is `NULL` if `messageTypes` is equal to
  `VK_DEBUG_UTILS_MESSAGE_TYPE_DEVICE_ADDRESS_BINDING_BIT_EXT`, or a
  null-terminated UTF-8 string detailing the trigger conditions.

- `queueLabelCount` is a count of items contained in the `pQueueLabels`
  array.

- `pQueueLabels` is `NULL` or a pointer to an array of
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) active in the
  current `VkQueue` at the time the callback was triggered. Refer to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debugging-queue-labels"
  target="_blank" rel="noopener">Queue Labels</a> for more information.

- `cmdBufLabelCount` is a count of items contained in the
  `pCmdBufLabels` array.

- `pCmdBufLabels` is `NULL` or a pointer to an array of
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) active in the
  current `VkCommandBuffer` at the time the callback was triggered.
  Refer to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debugging-command-buffer-labels"
  target="_blank" rel="noopener">Command Buffer Labels</a> for more
  information.

- `objectCount` is a count of items contained in the `pObjects` array.

- `pObjects` is a pointer to an array of
  [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html)
  objects related to the detected issue. The array is roughly in order
  or importance, but the 0th element is always guaranteed to be the most
  important object for this message.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This structure should only be considered valid during the lifetime of
the triggered callback.</p></td>
</tr>
</tbody>
</table>

Since adding queue and command buffer labels behaves like pushing and
popping onto a stack, the order of both `pQueueLabels` and
`pCmdBufLabels` is based on the order the labels were defined. The
result is that the first label in either `pQueueLabels` or
`pCmdBufLabels` will be the first defined (and therefore the oldest)
while the last label in each list will be the most recent.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>pQueueLabels</code> will only be non-<code>NULL</code> if one
of the objects in <code>pObjects</code> can be related directly to a
defined <code>VkQueue</code> which has had one or more labels associated
with it.</p>
<p>Likewise, <code>pCmdBufLabels</code> will only be
non-<code>NULL</code> if one of the objects in <code>pObjects</code> can
be related directly to a defined <code>VkCommandBuffer</code> which has
had one or more labels associated with it. Additionally, while command
buffer labels allow for beginning and ending across different command
buffers, the debug messaging framework <strong>cannot</strong> guarantee
that labels in <code>pCmdBufLables</code> will contain those defined
outside of the associated command buffer. This is partially due to the
fact that the association of one command buffer with another may not
have been defined at the time the debug message is triggered.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-sType-sType"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-sType-sType"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT`

- <a href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-pNext-pNext"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-pNext-pNext"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingCallbackDataEXT.html)

- <a href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-sType-unique"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-sType-unique"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-flags-zerobitmask"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-flags-zerobitmask"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-pMessageIdName-parameter"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-pMessageIdName-parameter"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-pMessageIdName-parameter  
  If `pMessageIdName` is not `NULL`, `pMessageIdName` **must** be a
  null-terminated UTF-8 string

- <a href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-pMessage-parameter"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-pMessage-parameter"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-pMessage-parameter  
  If `pMessage` is not `NULL`, `pMessage` **must** be a null-terminated
  UTF-8 string

- <a
  href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-pQueueLabels-parameter"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-pQueueLabels-parameter"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-pQueueLabels-parameter  
  If `queueLabelCount` is not `0`, `pQueueLabels` **must** be a valid
  pointer to an array of `queueLabelCount` valid
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structures

- <a
  href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-pCmdBufLabels-parameter"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-pCmdBufLabels-parameter"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-pCmdBufLabels-parameter  
  If `cmdBufLabelCount` is not `0`, `pCmdBufLabels` **must** be a valid
  pointer to an array of `cmdBufLabelCount` valid
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structures

- <a href="#VUID-VkDebugUtilsMessengerCallbackDataEXT-pObjects-parameter"
  id="VUID-VkDebugUtilsMessengerCallbackDataEXT-pObjects-parameter"></a>
  VUID-VkDebugUtilsMessengerCallbackDataEXT-pObjects-parameter  
  If `objectCount` is not `0`, `pObjects` **must** be a valid pointer to
  an array of `objectCount` valid
  [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkDebugUtilsMessengerCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDebugUtilsMessengerCallbackEXT.html),
[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html),
[VkDebugUtilsMessengerCallbackDataFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataFlagsEXT.html),
[VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSubmitDebugUtilsMessageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSubmitDebugUtilsMessageEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsMessengerCallbackDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

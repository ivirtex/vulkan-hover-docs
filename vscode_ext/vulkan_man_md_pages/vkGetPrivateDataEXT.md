# vkGetPrivateData(3) Manual Page

## Name

vkGetPrivateData - Retrieve data associated with a Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve application-defined data from a slot associated with a
Vulkan object, call:

``` c
// Provided by VK_VERSION_1_3
void vkGetPrivateData(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t*                                   pData);
```

or the equivalent command

``` c
// Provided by VK_EXT_private_data
void vkGetPrivateDataEXT(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t*                                   pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the object

- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) specifying the
  type of object data is associated with.

- `objectHandle` is a handle to the object data is associated with.

- `privateDataSlot` is a handle to a
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) specifying location of
  private data pointer storage.

- `pData` is a pointer to specify where application-defined data is
  returned. `0` will be written in the absence of a previous call to
  `vkSetPrivateData` using the object specified by `objectHandle`.

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
<p>Due to platform details on Android, implementations might not be able
to reliably return <code>0</code> from calls to
<code>vkGetPrivateData</code> for <a
href="VkSwapchainKHR.html">VkSwapchainKHR</a> objects on which
<code>vkSetPrivateData</code> has not previously been called. This
erratum is exclusive to the Android platform and objects of type <a
href="VkSwapchainKHR.html">VkSwapchainKHR</a>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkGetPrivateData-objectType-04018"
  id="VUID-vkGetPrivateData-objectType-04018"></a>
  VUID-vkGetPrivateData-objectType-04018  
  `objectHandle` **must** be `device` or a child of `device`

- <a href="#VUID-vkGetPrivateData-objectHandle-09498"
  id="VUID-vkGetPrivateData-objectHandle-09498"></a>
  VUID-vkGetPrivateData-objectHandle-09498  
  `objectHandle` **must** be a valid handle to an object of type
  `objectType`

Valid Usage (Implicit)

- <a href="#VUID-vkGetPrivateData-device-parameter"
  id="VUID-vkGetPrivateData-device-parameter"></a>
  VUID-vkGetPrivateData-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetPrivateData-objectType-parameter"
  id="VUID-vkGetPrivateData-objectType-parameter"></a>
  VUID-vkGetPrivateData-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html)
  value

- <a href="#VUID-vkGetPrivateData-privateDataSlot-parameter"
  id="VUID-vkGetPrivateData-privateDataSlot-parameter"></a>
  VUID-vkGetPrivateData-privateDataSlot-parameter  
  `privateDataSlot` **must** be a valid
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) handle

- <a href="#VUID-vkGetPrivateData-pData-parameter"
  id="VUID-vkGetPrivateData-pData-parameter"></a>
  VUID-vkGetPrivateData-pData-parameter  
  `pData` **must** be a valid pointer to a `uint64_t` value

- <a href="#VUID-vkGetPrivateData-privateDataSlot-parent"
  id="VUID-vkGetPrivateData-privateDataSlot-parent"></a>
  VUID-vkGetPrivateData-privateDataSlot-parent  
  `privateDataSlot` **must** have been created, allocated, or retrieved
  from `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_private_data](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_private_data.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html),
[VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPrivateData"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

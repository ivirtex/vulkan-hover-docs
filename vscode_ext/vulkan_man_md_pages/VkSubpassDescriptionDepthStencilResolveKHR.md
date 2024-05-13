# VkSubpassDescriptionDepthStencilResolve(3) Manual Page

## Name

VkSubpassDescriptionDepthStencilResolve - Structure specifying
depth/stencil resolve operations for a subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassDescriptionDepthStencilResolve` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSubpassDescriptionDepthStencilResolve {
    VkStructureType                  sType;
    const void*                      pNext;
    VkResolveModeFlagBits            depthResolveMode;
    VkResolveModeFlagBits            stencilResolveMode;
    const VkAttachmentReference2*    pDepthStencilResolveAttachment;
} VkSubpassDescriptionDepthStencilResolve;
```

or the equivalent

``` c
// Provided by VK_KHR_depth_stencil_resolve
typedef VkSubpassDescriptionDepthStencilResolve VkSubpassDescriptionDepthStencilResolveKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `depthResolveMode` is a
  [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html) value describing
  the depth resolve mode.

- `stencilResolveMode` is a
  [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html) value describing
  the stencil resolve mode.

- `pDepthStencilResolveAttachment` is `NULL` or a pointer to a
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure
  defining the depth/stencil resolve attachment for this subpass and its
  layout.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) includes a
`VkSubpassDescriptionDepthStencilResolve` structure, then that structure
describes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
target="_blank" rel="noopener">multisample resolve operations</a> for
the depth/stencil attachment in a subpass. If this structure is not
included in the `pNext` chain of
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html), or if it is and
either `pDepthStencilResolveAttachment` is `NULL` or its attachment
index is `VK_ATTACHMENT_UNUSED`, it indicates that no depth/stencil
resolve attachment will be used in the subpass.

Valid Usage

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03177"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03177"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03177  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, `pDepthStencilAttachment` **must**
  not be `NULL` or have the value `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03179"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03179"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03179  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, `pDepthStencilAttachment` **must**
  not have a sample count of `VK_SAMPLE_COUNT_1_BIT`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03180"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03180"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03180  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, `pDepthStencilResolveAttachment`
  **must** have a sample count of `VK_SAMPLE_COUNT_1_BIT`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-02651"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-02651"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-02651  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED` then it **must** have an image format
  whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> contain
  `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03181"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03181"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03181  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED` and [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `pDepthStencilResolveAttachment` has a depth component, then the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `pDepthStencilAttachment` **must** have a
  depth component with the same number of bits and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-numericformat"
  target="_blank" rel="noopener">numeric format</a>

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03182"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03182"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03182  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, and [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `pDepthStencilResolveAttachment` has a stencil component, then the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `pDepthStencilAttachment` **must** have a
  stencil component with the same number of bits and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-numericformat"
  target="_blank" rel="noopener">numeric format</a>

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03178"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03178"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03178  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, `depthResolveMode` and
  `stencilResolveMode` **must** not both be `VK_RESOLVE_MODE_NONE`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-depthResolveMode-03183"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-depthResolveMode-03183"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-depthResolveMode-03183  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED` and the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `pDepthStencilResolveAttachment` has a depth component, then the value
  of `depthResolveMode` **must** be one of the bits set in
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedDepthResolveModes`
  or `VK_RESOLVE_MODE_NONE`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-stencilResolveMode-03184"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-stencilResolveMode-03184"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-stencilResolveMode-03184  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED` and the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `pDepthStencilResolveAttachment` has a stencil component, then the
  value of `stencilResolveMode` **must** be one of the bits set in
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedStencilResolveModes`
  or `VK_RESOLVE_MODE_NONE`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03185"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03185"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03185  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `pDepthStencilResolveAttachment` has both depth and stencil
  components,
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve`
  is `VK_FALSE`, and
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone`
  is `VK_FALSE`, then the values of `depthResolveMode` and
  `stencilResolveMode` **must** be identical

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03186"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03186"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03186  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have
  the value `VK_ATTACHMENT_UNUSED`, the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `pDepthStencilResolveAttachment` has both depth and stencil
  components,
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve`
  is `VK_FALSE` and
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone`
  is `VK_TRUE`, then the values of `depthResolveMode` and
  `stencilResolveMode` **must** be identical or one of them **must** be
  `VK_RESOLVE_MODE_NONE`

- <a href="#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06873"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06873"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06873  
  If the `pNext` chain of
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) includes a
  `VkMultisampledRenderToSingleSampledInfoEXT` structure, the
  `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, and
  `pDepthStencilAttachment` is not `NULL` and does not have the value
  `VK_ATTACHMENT_UNUSED`, `depthResolveMode` and `stencilResolveMode`
  **must** not both be `VK_RESOLVE_MODE_NONE`

- <a href="#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06874"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06874"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06874  
  If the `pNext` chain of
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) includes a
  `VkMultisampledRenderToSingleSampledInfoEXT` structure whose
  `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, and
  `pDepthStencilAttachment` is not `NULL`, does not have the value
  `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a
  depth component, then the value of `depthResolveMode` **must** be one
  of the bits set in
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedDepthResolveModes`
  or `VK_RESOLVE_MODE_NONE`

- <a href="#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06875"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06875"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06875  
  If the `pNext` chain of
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) includes a
  `VkMultisampledRenderToSingleSampledInfoEXT` structure whose
  `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, and
  `pDepthStencilAttachment` is not `NULL`, does not have the value
  `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) with a
  stencil component, then the value of `stencilResolveMode` **must** be
  one of the bits set in
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedStencilResolveModes`
  or `VK_RESOLVE_MODE_NONE`

- <a href="#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06876"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06876"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06876  
  If the `pNext` chain of
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) includes a
  `VkMultisampledRenderToSingleSampledInfoEXT` structure whose
  `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`,
  `pDepthStencilAttachment` is not `NULL`, does not have the value
  `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) with both
  depth and stencil components, and both
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve`
  and
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone`
  are `VK_FALSE`, then the values of `depthResolveMode` and
  `stencilResolveMode` **must** be identical

- <a href="#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06877"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06877"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06877  
  If the `pNext` chain of
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) includes a
  `VkMultisampledRenderToSingleSampledInfoEXT` structure whose
  `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`,
  `pDepthStencilAttachment` is not `NULL`, does not have the value
  `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) with both
  depth and stencil components,
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve`
  is `VK_FALSE`, and
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone`
  is `VK_TRUE`, then the values of `depthResolveMode` and
  `stencilResolveMode` **must** be identical or one of them **must** be
  `VK_RESOLVE_MODE_NONE`

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassDescriptionDepthStencilResolve-sType-sType"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-sType-sType"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE`

- <a
  href="#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-parameter"
  id="VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-parameter"></a>
  VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-parameter  
  If `pDepthStencilResolveAttachment` is not `NULL`,
  `pDepthStencilResolveAttachment` **must** be a valid pointer to a
  valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_depth_stencil_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_depth_stencil_resolve.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html),
[VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassDescriptionDepthStencilResolve"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

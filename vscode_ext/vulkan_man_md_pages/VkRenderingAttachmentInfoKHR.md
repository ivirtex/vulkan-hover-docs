# VkRenderingAttachmentInfo(3) Manual Page

## Name

VkRenderingAttachmentInfo - Structure specifying attachment information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderingAttachmentInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkRenderingAttachmentInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkImageView              imageView;
    VkImageLayout            imageLayout;
    VkResolveModeFlagBits    resolveMode;
    VkImageView              resolveImageView;
    VkImageLayout            resolveImageLayout;
    VkAttachmentLoadOp       loadOp;
    VkAttachmentStoreOp      storeOp;
    VkClearValue             clearValue;
} VkRenderingAttachmentInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_dynamic_rendering
typedef VkRenderingAttachmentInfo VkRenderingAttachmentInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageView` is the image view that will be used for rendering.

- `imageLayout` is the layout that `imageView` will be in during
  rendering.

- `resolveMode` is a [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html)
  value defining how data written to `imageView` will be resolved into
  `resolveImageView`.

- `resolveImageView` is an image view used to write resolved data at the
  end of rendering.

- `resolveImageLayout` is the layout that `resolveImageView` will be in
  during rendering.

- `loadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value
  defining the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">load operation</a> for the attachment.

- `storeOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value
  defining the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
  target="_blank" rel="noopener">store operation</a> for the attachment.

- `clearValue` is a [VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html) structure defining
  values used to clear `imageView` when `loadOp` is
  `VK_ATTACHMENT_LOAD_OP_CLEAR`.

## <a href="#_description" class="anchor"></a>Description

Values in `imageView` are loaded and stored according to the values of
`loadOp` and `storeOp`, within the render area for each device specified
in [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html). If `imageView` is
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and `resolveMode` is not
`VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, other members of
this structure are ignored; writes to this attachment will be discarded,
and no <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
target="_blank" rel="noopener">load</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
target="_blank" rel="noopener">store</a>, or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
target="_blank" rel="noopener">multisample resolve</a> operations will
be performed.

If `resolveMode` is `VK_RESOLVE_MODE_NONE`, then `resolveImageView` is
ignored. If `resolveMode` is not `VK_RESOLVE_MODE_NONE`, and
`resolveImageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
target="_blank" rel="noopener">render pass multisample resolve
operation</a> is defined for the attachment subresource. If
`resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
and the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
target="_blank"
rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
limit is `VK_TRUE`, values are only undefined once <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
target="_blank" rel="noopener">load operations</a> have completed.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The resolve mode and store operation are independent; it is valid to
write both resolved and unresolved values, and equally valid to discard
the unresolved values while writing the resolved ones.</p></td>
</tr>
</tbody>
</table>

Store and resolve operations are only performed at the end of a render
pass instance that does not specify the
`VK_RENDERING_SUSPENDING_BIT_KHR` flag.

Load operations are only performed at the beginning of a render pass
instance that does not specify the `VK_RENDERING_RESUMING_BIT_KHR` flag.

Image contents at the end of a suspended render pass instance remain
defined for access by a resuming render pass instance.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
target="_blank"
rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
limit is `VK_TRUE`, and `resolveMode` is
`VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, values in the
color attachment will be loaded from the resolve attachment at the start
of rendering, and **may** also be reloaded any time after a resolve
occurs or the resolve attachment is written to; if this occurs it
**must** happen-before any writes to the color attachment are performed
which happen-after the resolve that triggers this. If any color
component in the external format is subsampled, values will be read from
the nearest sample in the image when they are loaded.

Valid Usage

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06129"
  id="VUID-VkRenderingAttachmentInfo-imageView-06129"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06129  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and has a
  non-integer color format, `resolveMode` **must** be
  `VK_RESOLVE_MODE_NONE` or `VK_RESOLVE_MODE_AVERAGE_BIT`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06130"
  id="VUID-VkRenderingAttachmentInfo-imageView-06130"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06130  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and has an
  integer color format, `resolveMode` **must** be `VK_RESOLVE_MODE_NONE`
  or `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06861"
  id="VUID-VkRenderingAttachmentInfo-imageView-06861"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06861  
  `imageView` **must** not have a sample count of
  `VK_SAMPLE_COUNT_1_BIT` if all of the following hold:

  - `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

  - `resolveMode` is not `VK_RESOLVE_MODE_NONE`

  - the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) does
    not include a
    [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
    structure with the `multisampledRenderToSingleSampledEnable` field
    equal to `VK_TRUE`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06862"
  id="VUID-VkRenderingAttachmentInfo-imageView-06862"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06862  
  `resolveImageView` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) if all of the following hold:

  - `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

  - `resolveMode` is not `VK_RESOLVE_MODE_NONE`

  - the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) does
    not include a
    [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
    structure with the `multisampledRenderToSingleSampledEnable` field
    equal to `VK_TRUE`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06863"
  id="VUID-VkRenderingAttachmentInfo-imageView-06863"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06863  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, the `pNext` chain of
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) includes a
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure with the `multisampledRenderToSingleSampledEnable` field
  equal to `VK_TRUE`, and `imageView` has a sample count of
  `VK_SAMPLE_COUNT_1_BIT`, `resolveImageView` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06864"
  id="VUID-VkRenderingAttachmentInfo-imageView-06864"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06864  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `resolveImageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageView`
  **must** have a sample count of `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06865"
  id="VUID-VkRenderingAttachmentInfo-imageView-06865"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06865  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `resolveImageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `imageView` and
  `resolveImageView` **must** have the same [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06135"
  id="VUID-VkRenderingAttachmentInfo-imageView-06135"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06135  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`,
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06136"
  id="VUID-VkRenderingAttachmentInfo-imageView-06136"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06136  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06137"
  id="VUID-VkRenderingAttachmentInfo-imageView-06137"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06137  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06138"
  id="VUID-VkRenderingAttachmentInfo-imageView-06138"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06138  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** not be
  `VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06139"
  id="VUID-VkRenderingAttachmentInfo-imageView-06139"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06139  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be `VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06140"
  id="VUID-VkRenderingAttachmentInfo-imageView-06140"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06140  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** not be
  `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06141"
  id="VUID-VkRenderingAttachmentInfo-imageView-06141"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06141  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06142"
  id="VUID-VkRenderingAttachmentInfo-imageView-06142"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06142  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06143"
  id="VUID-VkRenderingAttachmentInfo-imageView-06143"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06143  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** not be
  `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06144"
  id="VUID-VkRenderingAttachmentInfo-imageView-06144"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06144  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be
  `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06145"
  id="VUID-VkRenderingAttachmentInfo-imageView-06145"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06145  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** not be `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-06146"
  id="VUID-VkRenderingAttachmentInfo-imageView-06146"></a>
  VUID-VkRenderingAttachmentInfo-imageView-06146  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout`
  **must** not be `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`

- <a href="#VUID-VkRenderingAttachmentInfo-externalFormatResolve-09323"
  id="VUID-VkRenderingAttachmentInfo-externalFormatResolve-09323"></a>
  VUID-VkRenderingAttachmentInfo-externalFormatResolve-09323  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is not enabled, `resolveMode` **must** not be
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-09324"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-09324"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-09324  
  If `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
  `resolveImageView` **must** be a valid image view

- <a
  href="#VUID-VkRenderingAttachmentInfo-nullColorAttachmentWithExternalFormatResolve-09325"
  id="VUID-VkRenderingAttachmentInfo-nullColorAttachmentWithExternalFormatResolve-09325"></a>
  VUID-VkRenderingAttachmentInfo-nullColorAttachmentWithExternalFormatResolve-09325  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  property is `VK_TRUE` and `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
  `resolveImageView` **must** have been created with an image with a
  `samples` value of `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-09326"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-09326"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-09326  
  If `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
  `resolveImageView` **must** have been created with an external format
  specified by [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-09327"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-09327"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-09327  
  If `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
  `resolveImageView` **must** have been created with a
  `subresourceRange.layerCount` of `1`

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-09328"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-09328"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-09328  
  If `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  is `VK_TRUE`, `imageView` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-09329"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-09329"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-09329  
  If `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  is `VK_FALSE`, `imageView` **must** be a valid
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-09330"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-09330"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-09330  
  If `resolveMode` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  is `VK_FALSE`, `imageView` **must** have a format equal to the value
  of
  [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)::`colorAttachmentFormat`
  as returned by a call to
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
  for the Android hardware buffer that was used to create
  `resolveImageView`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderingAttachmentInfo-sType-sType"
  id="VUID-VkRenderingAttachmentInfo-sType-sType"></a>
  VUID-VkRenderingAttachmentInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO`

- <a href="#VUID-VkRenderingAttachmentInfo-pNext-pNext"
  id="VUID-VkRenderingAttachmentInfo-pNext-pNext"></a>
  VUID-VkRenderingAttachmentInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkRenderingAttachmentInfo-imageView-parameter"
  id="VUID-VkRenderingAttachmentInfo-imageView-parameter"></a>
  VUID-VkRenderingAttachmentInfo-imageView-parameter  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-VkRenderingAttachmentInfo-imageLayout-parameter"
  id="VUID-VkRenderingAttachmentInfo-imageLayout-parameter"></a>
  VUID-VkRenderingAttachmentInfo-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-VkRenderingAttachmentInfo-resolveMode-parameter"
  id="VUID-VkRenderingAttachmentInfo-resolveMode-parameter"></a>
  VUID-VkRenderingAttachmentInfo-resolveMode-parameter  
  If `resolveMode` is not `0`, `resolveMode` **must** be a valid
  [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html) value

- <a href="#VUID-VkRenderingAttachmentInfo-resolveImageView-parameter"
  id="VUID-VkRenderingAttachmentInfo-resolveImageView-parameter"></a>
  VUID-VkRenderingAttachmentInfo-resolveImageView-parameter  
  If `resolveImageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `resolveImageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  handle

- <a href="#VUID-VkRenderingAttachmentInfo-resolveImageLayout-parameter"
  id="VUID-VkRenderingAttachmentInfo-resolveImageLayout-parameter"></a>
  VUID-VkRenderingAttachmentInfo-resolveImageLayout-parameter  
  `resolveImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkRenderingAttachmentInfo-loadOp-parameter"
  id="VUID-VkRenderingAttachmentInfo-loadOp-parameter"></a>
  VUID-VkRenderingAttachmentInfo-loadOp-parameter  
  `loadOp` **must** be a valid
  [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value

- <a href="#VUID-VkRenderingAttachmentInfo-storeOp-parameter"
  id="VUID-VkRenderingAttachmentInfo-storeOp-parameter"></a>
  VUID-VkRenderingAttachmentInfo-storeOp-parameter  
  `storeOp` **must** be a valid
  [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value

- <a href="#VUID-VkRenderingAttachmentInfo-clearValue-parameter"
  id="VUID-VkRenderingAttachmentInfo-clearValue-parameter"></a>
  VUID-VkRenderingAttachmentInfo-clearValue-parameter  
  `clearValue` **must** be a valid [VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html)
  union

- <a href="#VUID-VkRenderingAttachmentInfo-commonparent"
  id="VUID-VkRenderingAttachmentInfo-commonparent"></a>
  VUID-VkRenderingAttachmentInfo-commonparent  
  Both of `imageView`, and `resolveImageView` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html),
[VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html),
[VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html),
[VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingAttachmentInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
